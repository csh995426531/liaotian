package test

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"liaotian/domain/message/proto"
	"net/http"
	"testing"
)

func sub(t *testing.T) {
	testData := []struct {
		UserId int64
		Code   int32
		Msg    string
		Ok     bool
	}{
		{1, http.StatusOK, "success", true},
		{0, http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewMessageService("domain.message.service", client.DefaultClient)

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			request := &proto.SubRequest{
				UserId: data.UserId,
			}

			resp, err := service.Sub(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if resp.Ok != data.Ok {
					t.Errorf("响应Ok错误, want:%v, got:%v", data.Ok, resp.Ok)
				}
			}
		})
	}
}

func unSub(t *testing.T) {
	testData := []struct {
		UserId int64
		Code   int32
		Msg    string
		Ok     bool
	}{
		{1, http.StatusOK, "success", true},
		{0, http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewMessageService("domain.message.service", client.DefaultClient)

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			request := &proto.UnSubRequest{
				UserId: data.UserId,
			}

			resp, err := service.UnSub(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if resp.Ok != data.Ok {
					t.Errorf("响应Ok错误, want:%v, got:%v", data.Ok, resp.Ok)
				}
			}
		})
	}
}

func send(t *testing.T) {
	testData := []struct {
		FriendId   int64
		SenderId   int64
		ReceiverId int64
		Content    string
		Code       int32
		Msg        string
		Ok         bool
	}{
		{1, 1, 2, "你好啊", http.StatusOK, "success", true},
		{0, 1, 2, "我是赛利亚", http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewMessageService("domain.message.service", client.DefaultClient)

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			request := &proto.SendRequest{
				FriendId:   data.FriendId,
				SenderId:   data.SenderId,
				ReceiverId: data.ReceiverId,
				Content:    data.Content,
			}

			resp, err := service.Send(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if resp.Ok != data.Ok {
					t.Errorf("响应Ok错误, want:%v, got:%v", data.Ok, resp.Ok)
				}
			}
		})
	}
}
