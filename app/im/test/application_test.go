package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"liaotian/app/im/handler"
	"liaotian/domain/friend/proto"
	"net/http"
	"testing"
)

func createApplication(t *testing.T) {

	testData := []struct {
		Token      string
		ReceiverId int64
		HttpCode   int
		Response   *proto.ApplicationResponse
	}{
		{"我是万能钥匙", 2, http.StatusCreated, &proto.ApplicationResponse{
			Data:    &proto.Application{Id: 1, SenderId: 1, ReceiverId: 2, SayList: []*proto.ApplicationSay{}},
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

func applicationList(t *testing.T) {
	testData := []struct {
		Token    string
		HttpCode int
		Response *proto.GetApplicationListResponse
	}{
		{"我是万能钥匙", http.StatusOK, &proto.GetApplicationListResponse{
			Data:    []*proto.Application{{Id: 1, SenderId: 1, ReceiverId: 2, SayList: []*proto.ApplicationSay{}}},
			Message: "success",
		}},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().GetApplicationList(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			client := &http.Client{}
			res, _ := http.NewRequest("GET", "http://127.0.0.1:18282/application/list", nil)
			res.Header.Set("token", data.Token)
			resp, err := client.Do(res)
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

func applicationInfo(t *testing.T) {
	testData := []struct {
		Token    string
		Id       int64
		HttpCode int
		Response *proto.ApplicationResponse
	}{
		{"我是万能钥匙", 1, http.StatusOK, &proto.ApplicationResponse{
			Data: &proto.Application{Id: 1, SenderId: 1, ReceiverId: 2, SayList: []*proto.ApplicationSay{{
				Id: 1, SenderId: 1, Content: "加一下好友",
			}}},
			Message: "success",
		}},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().GetApplicationInfo(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			client := &http.Client{}
			url := fmt.Sprintf("http://127.0.0.1:18282/application/info?id=%v", data.Id)
			res, _ := http.NewRequest("GET", url, nil)
			res.Header.Set("token", data.Token)
			resp, err := client.Do(res)
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

func passApplication(t *testing.T) {

	testData := []struct {
		Token    string
		Id       int64
		HttpCode int
		Response *proto.Response
	}{
		{"我是万能钥匙", 1, http.StatusOK, &proto.Response{
			Ok:      true,
			Message: "success",
		}},
	}

	for _, data := range testData {

		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().PassApplicationInfo(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			bytesData, _ := json.Marshal(data)
			reader := bytes.NewReader(bytesData)
			resp, err := http.Post("http://127.0.0.1:18282/application/pass", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}

			gotBody, _ := ioutil.ReadAll(resp.Body)
			wantBody, _ := json.Marshal(gin.H{
				"msg":  "success",
				"data": data.Response.Ok,
			})

			if string(gotBody) != string(wantBody) {
				t.Errorf("响应body错误, want:%v, got:%v", string(wantBody), string(gotBody))
			}
		})
	}
}

func rejectApplication(t *testing.T) {

	testData := []struct {
		Token    string
		Id       int64
		HttpCode int
		Response *proto.Response
	}{
		{"我是万能钥匙", 1, http.StatusOK, &proto.Response{
			Ok:      true,
			Message: "success",
		}},
	}

	for _, data := range testData {

		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().RejectApplicationInfo(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			bytesData, _ := json.Marshal(data)
			reader := bytes.NewReader(bytesData)
			resp, err := http.Post("http://127.0.0.1:18282/application/reject", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}

			gotBody, _ := ioutil.ReadAll(resp.Body)
			wantBody, _ := json.Marshal(gin.H{
				"msg":  "success",
				"data": data.Response.Ok,
			})

			if string(gotBody) != string(wantBody) {
				t.Errorf("响应body错误, want:%v, got:%v", string(wantBody), string(gotBody))
			}
		})
	}
}

func replyApplication(t *testing.T) {
	testData := []struct {
		Token    string
		Id       int64
		SenderId int64
		Content  string
		HttpCode int
		Response *proto.CreateApplicationSayResponse
	}{
		{"我是万能钥匙", 1, 1, "加我吧", http.StatusOK, &proto.CreateApplicationSayResponse{
			Data:    &proto.ApplicationSay{Id: 1, SenderId: 1, Content: "加我吧"},
			Message: "success",
		}},
	}

	for _, data := range testData {

		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().CreateApplicationSay(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			bytesData, _ := json.Marshal(data)
			reader := bytes.NewReader(bytesData)
			resp, err := http.Post("http://127.0.0.1:18282/application/reply", "application/json", reader)
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
