package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"liaotian/app/im/handler"
	"liaotian/domain/friend/proto"
	"net/http"
	"testing"
)

func CreateApplication(t *testing.T) {

	testData := []struct{
		Token   string
		ReceiverId int64
		HttpCode int
		Response *proto.ApplicationResponse
	}{
		{"我是万能钥匙",
			2, http.StatusCreated, &proto.ApplicationResponse{
			Data: &proto.Application{Id: 1, SenderId: 1, ReceiverId: 2, SayList: []*proto.ApplicationSay{}},
			Message: "success",
		}},
	}

	for _, data := range testData {

		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().CreateApplicationInfo(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			bytesData, _ := json.Marshal(data)
			reader := bytes.NewReader(bytesData)
			resp, err := http.Post("http://127.0.0.1:18282/application/create", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}

			gotBody, _ := ioutil.ReadAll(resp.Body)
			wantBody, _ := json.Marshal(gin.H{
				"msg":  "success",
				"data": data.Response.Data,
			})

			if string(gotBody) != string(wantBody) {
				t.Errorf("响应body错误, want:%v, got:%v", string(wantBody), string(gotBody))
			}
		})
	}
}
