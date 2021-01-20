package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"liaotian/app/im/handler"
	"liaotian/domain/friend/proto"
	userService "liaotian/domain/user/proto"
	"net/http"
	"testing"
)

func friendList(t *testing.T) {
	testData := []struct {
		Token    string
		HttpCode int
		Response *proto.GetFriendListResponse
	}{
		{"我是万能钥匙", http.StatusOK, &proto.GetFriendListResponse{
			Message: "success",
			Data:    []*proto.FriendList{{Id: 1, UserId: 2}},
		}},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().GetFriendList(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			client := &http.Client{}
			res, _ := http.NewRequest("GET", "http://127.0.0.1:18282/friend/list", nil)
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
				"msg": "success",
				"data": []struct {
					Id      int64  `json:"id"`
					UserId  int64  `json:"user_id"`
					Name    string `json:"name"`
					Avatar  string `json:"avatar"`
					Account string `json:"account"`
				}{
					{Id: 1, UserId: 2, Name: "name", Avatar: "avatar", Account: "account"},
				},
			})
			if string(gotBody) != string(wantBody) {
				t.Errorf("响应body错误, want:%v, got:%v", string(wantBody), string(gotBody))
			}
		})
	}
}

func deleteFriendInfo(t *testing.T) {
	testData := []struct {
		Token    string
		Id       int64
		HttpCode int
		Response *proto.Response
	}{
		{"我是万能钥匙", 1, http.StatusOK, &proto.Response{
			Message: "success",
			Ok:      true,
		}},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service := NewMockFriendService(ctrl)
			handler.FriendDomain(service)
			service.EXPECT().DeleteFriendInfo(gomock.Any(), gomock.Any()).Return(data.Response, nil)

			client := &http.Client{}
			url := fmt.Sprintf("http://127.0.0.1:18282/friend/delete?id=%v", data.Id)
			res, _ := http.NewRequest("DELETE", url, nil)
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
				"data": data.Response.Ok,
			})
			if string(gotBody) != string(wantBody) {
				t.Errorf("响应body错误, want:%v, got:%v", string(wantBody), string(gotBody))
			}
		})
	}
}

func friendInfo(t *testing.T) {
	testData := []struct {
		Token    string
		Id       int64
		HttpCode int
		Response *userService.Response
	}{
		{"我是万能钥匙", 1, http.StatusOK, &userService.Response{
			Message: "success",
			Data:    &userService.User{Id: 1, Name: "张三", Account: "zhangsan", Avatar: "http://www.baidu.com"},
		}},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			client := &http.Client{}
			url := fmt.Sprintf("http://127.0.0.1:18282/friend/info?id=%v", data.Id)
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
