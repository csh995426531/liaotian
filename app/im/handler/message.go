package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"liaotian/app/im/event"
	"liaotian/app/im/handler/validator"
	"liaotian/domain/message/proto"
	ginResult "liaotian/middlewares/common-result/gin"
	"liaotian/middlewares/logger/zap"
	"net/http"
)

/**
消息应用服务
*/

//连接
func Connect(ctx *gin.Context) {
	connRequestValidator := &validator.ConnRequest{}
	req := &proto.SubRequest{}
	if err := validator.Bind(ctx, connRequestValidator, req); err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	wsSocket, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		zap.SugarLogger.Errorf("upGrader.Upgrade error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "连接异常")
		return
	}

	if res, err := DomainMessage.Sub(ctx.Request.Context(), req); err != nil || res.Ok != true {
		zap.SugarLogger.Errorf("DomainMessage.Sub error: %v, res:%v", err, res)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		Close(wsSocket, req.UserId)
		return
	}

	//启动一个读协程，将数据推送到消息领域服务
	go func(wsSocket *websocket.Conn) {
		for {
			_, data, err := wsSocket.ReadMessage()
			if err != nil {
				Close(wsSocket, req.UserId)
				break
			}
			if len(data) == 0 {
				continue
			}
			sendRequestValidator := &validator.SendRequest{}
			if err := json.Unmarshal(data, sendRequestValidator); err != nil {
				if err := wsSocket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("数据格式错误,%v", err))); err != nil {
					zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
					Close(wsSocket, req.UserId)
					panic(err)
				}
				continue
			}

			sendRequestValidator.SenderId = req.UserId
			sendReq := &proto.SendRequest{}
			if err := validator.ExecBind(sendRequestValidator, sendReq); err != nil {
				if err := wsSocket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("数据格式错误,%v", err))); err != nil {
					zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
					Close(wsSocket, req.UserId)
					panic(err)
				}
				continue
			}
			if res, err := DomainMessage.Send(ctx.Request.Context(), sendReq); err != nil || !res.Ok {
				zap.SugarLogger.Errorf("DomainMessage.Send error: %v, res:%v", err, res)
				Close(wsSocket, req.UserId)
				panic(err)
			}
			if err := wsSocket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("send ok! {%v}", sendReq.String()))); err != nil {
				zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
				Close(wsSocket, req.UserId)
				panic(err)
			}
		}
	}(wsSocket)

	// 启动一个写协程，从消息领域服务接收消息
	go func(wsSocket *websocket.Conn) {
		if err := wsSocket.WriteMessage(websocket.TextMessage, []byte("连接成功")); err != nil {
			zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
			Close(wsSocket, req.UserId)
			panic(err)
		}
		for {
			data := event.Instance.ReadNewMessage(req.UserId)
			if err := wsSocket.WriteMessage(websocket.TextMessage, data); err != nil {
				zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
				Close(wsSocket, req.UserId)
				panic(err)
			}
		}
	}(wsSocket)
}

func Close(wsSocket *websocket.Conn, UserId int64) {
	_ = wsSocket.Close()
	if res, err := DomainMessage.UnSub(context.Background(), &proto.UnSubRequest{
		UserId: UserId,
	}); err != nil || !res.Ok {
		zap.SugarLogger.Errorf("DomainMessage.UnSub error: %v, res: %v", err, res)
	}
}
