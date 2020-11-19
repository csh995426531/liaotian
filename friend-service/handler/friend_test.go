package handler

import (
	"context"
	"fmt"
	"liaotian/friend-service/config"
	proto "liaotian/friend-service/proto/friend"
	"liaotian/friend-service/repository"
	"strconv"
	"testing"
)

var (
	handler *Handler
)

func TestMain(m *testing.M) {
	config.MysqlConfig.Url = "debian-sys-maint:F0sm3f7WrNJox1oV@tcp(129.211.55.205:3306)/liaotian"
	handler = New(repository.Init())
	m.Run()
}

func TestHandler_Add(t *testing.T) {

	addTests := []struct {
		OperatorId int64
		BuddyId    int64
		Resp       *proto.Response
	}{
		{1, 2, &proto.Response{Code: 200, Message: "SUCCESS"}},
		{1, 3, &proto.Response{Code: 200, Message: "SUCCESS"}},
		{1, 4, &proto.Response{Code: 200, Message: "SUCCESS"}},
		{1, 5, &proto.Response{Code: 200, Message: "SUCCESS"}},
		{1, 6, &proto.Response{Code: 200, Message: "SUCCESS"}},
	}

	for i, test := range addTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			resp := proto.Response{}
			err := handler.Add(context.Background(), &proto.AddRequest{OperatorId: test.OperatorId, BuddyId: test.BuddyId}, &resp)
			if err != nil {
				t.Error(err)
			}

			expected := fmt.Sprint(test.Resp)
			got := fmt.Sprint(resp)
			t.Logf("expected=%s, got=%s", expected, got)
			if resp.Code != test.Resp.Code {
				t.Errorf("add failed to input %+v, expected %s , got %s", test, expected, got)
			}
		})
	}
}

func TestHandler_List(t *testing.T) {

	listTests := []struct {
		OperatorId int64
		Offset     int64
		Limit      int64
		Resp       *proto.Response
	}{
		{1, 0, 2, &proto.Response{Code: 200, Message: "SUCCESS"}},
		{1, 2, 2, &proto.Response{Code: 200, Message: "SUCCESS"}},
	}

	for _, test := range listTests {
		t.Run("", func(t *testing.T) {
			resp := proto.Response{}
			err := handler.List(context.Background(), &proto.ListRequest{OperatorId: test.OperatorId, Offset: test.Offset, Limit: test.Limit}, &resp)
			if err != nil {
				t.Error(err)
			}

			expected := fmt.Sprint(test.Resp)
			got := fmt.Sprint(resp)
			if resp.Code != test.Resp.Code {
				t.Errorf("list failed to input: %+v, expected:%s, got:%s", test, expected, got)
			} else {

				for _, temp := range resp.List {
					getResp := proto.Response{}
					err := handler.Get(context.Background(), &proto.Request{FriendId: temp.Id}, &getResp)
					if err != nil {
						t.Error(err)
					}
					got := fmt.Sprint(resp)
					if getResp.Code != test.Resp.Code {
						t.Errorf("get failed to input: %+v, got:%s", temp, got)
					}

					delResp := proto.Response{}
					err = handler.Del(context.Background(), &proto.Request{FriendId: temp.Id}, &delResp)
					if err != nil {
						t.Error(err)
					}

					got = fmt.Sprint(delResp)
					if delResp.Code != test.Resp.Code {
						t.Errorf("del failed to input： %+v, got:%s", temp, got)
					}
				}
			}
		})
	}
}
