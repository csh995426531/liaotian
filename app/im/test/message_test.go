package test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"liaotian/app/im/handler"
	"liaotian/domain/message/proto"
	"net/http"
	"net/url"
	"testing"
)

func connect(t *testing.T) {
	type messageType struct {
		Id         int64  `json:"id"`
		ReceiverId int64  `json:"receiver_id"`
		Content    string `json:"content"`
	}
	testData := []struct {
		Token       string
		SendContent []messageType
		HttpCode    int
		Response    *proto.Response
	}{
		{"我是万能钥匙", []messageType{
			{1, 2, "你好啊"},
			{1, 2, "我是赛利亚"},
		}, http.StatusOK, &proto.Response{
			Ok:      true,
			Message: "success",
		}},
	}
	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockMessageService(ctrl)
			handler.MessageDomain(service)
			service.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			urlStr := url.URL{Scheme: "ws", Host: "127.0.0.1:18282", Path: "/message/connect"}
			var (
				dialer *websocket.Dialer
				client *websocket.Conn
			)
			header := http.Header{}
			header.Set("token", data.Token)
			client, _, err := dialer.Dial(urlStr.String(), header)
			if err != nil {
				t.Errorf("websocket Dial error: %v", err)
				panic(err)
			}

			for _, message := range data.SendContent {
				service.EXPECT().Send(gomock.Any(), gomock.Any()).Return(data.Response, nil)

				messageByte, _ := json.Marshal(message)
				err := client.WriteMessage(websocket.TextMessage, messageByte)
				if err != nil {
					t.Errorf("client.WriteMessage error:%v", err)
				}
			}

			for {
				_, msg, err := client.ReadMessage()
				if err != nil {
					t.Error(err)
				}
				if len(msg) == 0 {
					break
				}
				fmt.Printf("ReadMessage:%v\n", string(msg))
			}
		})
	}
}
