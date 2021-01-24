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
	msgSocket "liaotian/middlewares/websocket"
	"net/http"
)

/**
消息应用服务
*/

type Coon struct {
	id       int64
	connect  *msgSocket.Connect
	isClosed bool
}

//连接
func Connect(ctx *gin.Context) {
	connRequestValidator := &validator.ConnRequest{}
	req := &proto.SubRequest{}
	if err := validator.Bind(ctx, connRequestValidator, req); err != nil {
		ginResult.Failed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	connect, err := msgSocket.New(ctx)
	if err != nil {
		zap.SugarLogger.Errorf("upGrader.Upgrade error: %v", err)
		ginResult.Failed(ctx, http.StatusInternalServerError, "连接异常")
		return
	}

	if res, err := DomainMessage.Sub(ctx.Request.Context(), req); err != nil || res.Ok != true {
		zap.SugarLogger.Errorf("DomainMessage.Sub error: %v, res:%v", err, res)
		ginResult.Failed(ctx, http.StatusInternalServerError, "上游服务异常")
		connect.Close()
		return
	}

	coon := &Coon{
		id:       req.UserId,
		connect:  connect,
		isClosed: false,
	}

	if err := coon.connect.Write(websocket.TextMessage, []byte("连接成功")); err != nil {
		coon.close()
		return
	}

	//启动一个读协程，将数据推送到消息领域服务
	go coon.readWorker(ctx)

	// 启动一个写协程，从消息领域服务接收消息
	go coon.writeWorker()
}

func (c *Coon) readWorker(ctx *gin.Context) {
	for {
		data, err := c.connect.Read()
		if err != nil {
			goto ERR
		}
		sendRequestValidator := &validator.SendRequest{}
		if err := json.Unmarshal(data.Data, sendRequestValidator); err != nil {
			if err := c.connect.Write(websocket.TextMessage, []byte(fmt.Sprintf("数据格式错误,%v", err))); err != nil {
				zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
				goto ERR
			}
			continue
		}

		sendRequestValidator.SenderId = c.id
		sendReq := &proto.SendRequest{}
		if err := validator.ExecBind(sendRequestValidator, sendReq); err != nil {
			if err := c.connect.Write(websocket.TextMessage, []byte(fmt.Sprintf("数据格式错误,%v", err))); err != nil {
				zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
				goto ERR
			}
			continue
		}
		if res, err := DomainMessage.Send(ctx.Request.Context(), sendReq); err != nil || !res.Ok {
			zap.SugarLogger.Errorf("DomainMessage.Send error: %v, res:%v", err, res)
			goto ERR
		}
		if err := c.connect.Write(websocket.TextMessage, []byte(fmt.Sprintf("send ok! {%v}", string(data.Data)))); err != nil {
			zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
			goto ERR
		}
	}
ERR:
	c.close()
}

func (c *Coon) writeWorker() {
	for {
		data := event.Instance.ReadNewMessage(c.id)
		if err := c.connect.Write(websocket.TextMessage, data); err != nil {
			zap.SugarLogger.Errorf("wsSocket.WriteMessage error: %v", err)
			goto ERR
		}
	}
ERR:
	c.close()
}

func (c *Coon) close() {
	c.connect.Close()
	if !c.isClosed {
		if res, err := DomainMessage.UnSub(context.Background(), &proto.UnSubRequest{
			UserId: c.id,
		}); err != nil || !res.Ok {
			zap.SugarLogger.Errorf("DomainMessage.UnSub error: %v, res: %v", err, res)
			return
		}
		c.isClosed = true
	}
}
