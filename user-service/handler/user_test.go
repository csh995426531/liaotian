package handler

import (
	"context"
	"fmt"
	"liaotian/basic/config"
	user "liaotian/user-service/proto/user"
	"liaotian/user-service/repository"
	"testing"
)

var (
	handler *Handler
	userModel	*repository.ModelUser
)

func TestMain(m *testing.M) {
	config.Init(func(options *config.Options) {
		options.Path = "../"
	})
	handler = New(repository.Init())
	m.Run()
}

func TestRun(t *testing.T) {
	t.Run("testCreate", testCreate)
	t.Run("testGet", testGet)
}

func testCreate(t *testing.T) {

	createTests := []struct{
		Name 		string
		Password	string
		Resp		*user.Response
	} {
		{"张三", "a123123123",&user.Response{Code: 200, Message: "SUCCESS"}},
		{"李四", "a123123123",&user.Response{Code: 200, Message: "SUCCESS"}},
	}

	for _, test := range createTests {

		t.Run(test.Name, func(t *testing.T) {
			resp := &user.Response{}
			err := handler.Create(context.Background(), &user.CreateRequest{
				Name: test.Name,
				Password: test.Password,
			}, resp)

			if err != nil {
				t.Error("failed to connect server  ", err)
			}

			expected := fmt.Sprint(test.Resp)
			got := fmt.Sprint(resp)
			t.Logf("expected=%s, got=%s", expected, got)
			if resp.Code != test.Resp.Code {
				t.Errorf("post failed to input %s, expected %v , got %s", test, expected, got)
			}
		})
	}
}

func testGet(t *testing.T) {

	getTests := []struct{
		Name 		string
		Password	string
		Resp 		*user.Response
	} {
		{"张三","a123123123",&user.Response{Code: 200, Message: "SUCCESS"}},
		{"张四","a123123123",&user.Response{Code: 200, Message: "SUCCESS"}},
	}

	for _, test := range getTests {
		t.Run(test.Name, func(t *testing.T) {
			resp := &user.Response{}
			err := handler.Get(context.Background(), &user.Request{
				Name: test.Name,
				Password: test.Password,
			}, resp)

			if err != nil {
				t.Error("failed to connect server  ", err)
			}

			expected := fmt.Sprint(test.Resp)
			got := fmt.Sprint(resp)
			t.Logf("expected=%s, got=%s", expected, got)
			if test.Resp.Code != resp.Code {
				t.Errorf("get failed to input %s, expected %v , got %s", test, expected, got)
			}
		})
	}
}