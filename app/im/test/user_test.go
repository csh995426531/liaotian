package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	client "github.com/micro/go-micro/client"
	"io/ioutil"
	"liaotian/app/im/handler"
	authService "liaotian/domain/auth/proto"
	userService "liaotian/domain/user/proto"
	"net/http"
	"testing"
)

type testService struct {
}

func (c *testService) CreateUserInfo(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)

	if in.Account == "" || in.Password == "" || in.Name == "" {
		out.Code = http.StatusBadRequest
		out.Message = "缺少参数！"
		out.Data = nil
		return out, nil
	}

	if in.Account == "zhangsan" {
		out.Code = http.StatusForbidden
		out.Message = "账户已被注册！"
		out.Data = nil
		return out, nil
	}

	out.Code = http.StatusCreated
	out.Message = "success"
	out.Data = &userService.User{
		Id:      1,
		Name:    in.Name,
		Account: in.Account,
		Avatar:  in.Avatar,
	}

	return out, nil
}
func (c *testService) GetUserInfo(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)
	if in.Account == "" && in.Name == "" && in.Id == 0 {
		out.Code = http.StatusBadRequest
		out.Message = "缺少参数！"
		out.Data = nil
		return out, nil
	}

	if in.Id == 1 || in.Account == "zhangsan" || in.Name == "张三" {
		out.Code = http.StatusOK
		out.Message = "success"
		out.Data = &userService.User{
			Id:      1,
			Name:    "张三",
			Account: "zhangsan",
			Avatar:  "http://www.baidu.com",
		}
	} else {
		out.Code = http.StatusNotFound
		out.Message = "用户不存在"
	}
	return out, nil
}
func (c *testService) UpdateUserInfo(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)
	if in.Id != 1 {
		out.Code = http.StatusNotFound
		out.Message = "用户不存在"
		out.Data = nil
		return out, nil
	}

	out.Code = http.StatusOK
	out.Message = "success"
	out.Data = &userService.User{
		Id:      1,
		Name:    in.Name,
		Account: in.Account,
		Avatar:  in.Avatar,
	}
	return out, nil
}
func (c *testService) CheckUserPwd(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)
	if in.Account != "zhangsan" {
		out.Code = http.StatusNotFound
		out.Message = "用户不存在"
		out.Data = nil
		return out, nil
	}

	if in.Password != "123456" {
		out.Code = http.StatusUnauthorized
		out.Message = "密码错误"
		out.Data = nil
		return out, nil
	}

	out.Code = http.StatusOK
	out.Message = "success"
	out.Data = &userService.User{
		Id:      1,
		Name:    "张三",
		Account: "zhangsan",
		Avatar:  "http://www.baidu.com",
	}

	return out, nil
}
func (c *testService) BatchGetUserInfo(ctx context.Context, in *userService.BatchGetUserInfoRequest, opts ...client.CallOption) (*userService.BatchGetUserInfoResponse, error) {

	out := new(userService.BatchGetUserInfoResponse)
	if len(in.Ids) == 0 {
		out.Code = http.StatusBadRequest
		out.Message = "参数错误"
		out.Data = nil
		return out, nil
	}

	out.Code = http.StatusOK
	out.Message = "success"
	out.Data = make([]*userService.User, 0)
	for _, id := range in.Ids {
		out.Data = append(out.Data, &userService.User{
			Id:      id,
			Name:    "name",
			Account: "account",
			Avatar:  "avatar",
		})
	}
	return out, nil
}

func register(t *testing.T) {

	testData := []struct {
		Account  string
		Password string
		Name     string
		Avatar   string
		HttpCode int
		Response string
	}{
		{"zhangsan", "123456", "张三", "http://baidu.com", http.StatusForbidden, "{\"data\":null,\"msg\":\"账户已被注册！\"}"},
		{"", "123456", "李四", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Account为必填字段\"}"},
		{"lisi", "", "李四", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password为必填字段\"}"},
		{"lisi", "123456", "", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name为必填字段\"}"},
		{"lisi", "123456", "李四", "", http.StatusCreated, "{\"data\":{\"id\":1,\"name\":\"李四\",\"account\":\"lisi\"},\"msg\":\"success\"}"},
		{"aaaaaaaaaaaaaaaaaaaaa", "123456", "李四", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Account长度不能超过20个字符\"}"},
		{"lisi", "aaaaaaaaaaaaaaaaaaaaa", "李四", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password长度不能超过20个字符\"}"},
		{"lisi", "123456", "aaaaaaaaaaaaaaaaaaaaa", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name长度不能超过20个字符\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			bytesData, _ := json.Marshal(data)
			reader := bytes.NewReader(bytesData)

			resp, err := http.Post("http://127.0.0.1:18282/user/register", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != data.Response {
				t.Errorf("响应body错误，want:%v, got:%v", data.Response, string(body))
			}
		})
	}
}

func login(t *testing.T) {
	testData := []struct {
		Account  string
		Password string
		HttpCode int
		Response string
	}{
		{"zhangsan", "123456", http.StatusOK, "{\"data\":{\"info\":{\"id\":1,\"name\":\"张三\",\"account\":\"zhangsan\",\"avatar\":\"http://www.baidu.com\"},\"token\":\"我是万能钥匙\"},\"msg\":\"success\"}"},
		{"zhangsan", "111111", http.StatusUnauthorized, "{\"data\":null,\"msg\":\"密码错误\"}"},
		{"aaa", "111111", http.StatusNotFound, "{\"data\":null,\"msg\":\"用户不存在\"}"},
		{"", "111111", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Account为必填字段\"}"},
		{"aaa", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password为必填字段\"}"},
		{"aaaaaaaaaaaaaaaaaaaaa", "123456", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Account长度不能超过20个字符\"}"},
		{"zhangsan", "aaaaaaaaaaaaaaaaaaaaa", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password长度不能超过20个字符\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			byteData, _ := json.Marshal(data)
			reader := bytes.NewReader(byteData)
			resp, err := http.Post("http://127.0.0.1:18282/user/login", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != data.Response {
				t.Errorf("响应body错误，want:%v, got:%v", data.Response, string(body))
			}
		})
	}
}

func getUserInfo(t *testing.T) {
	testData := []struct {
		Token    string
		HttpCode int
		Response string
	}{
		{"我是万能钥匙", http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"张三\",\"account\":\"zhangsan\",\"avatar\":\"http://www.baicu.com\"},\"msg\":\"success\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			client := &http.Client{}
			res, _ := http.NewRequest("GET", "http://127.0.0.1:18282/user/info", nil)
			res.Header.Set("token", data.Token)
			resp, err := client.Do(res)
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != data.Response {
				t.Errorf("响应body错误，want:%v, got: %v", data.Response, string(body))
			}
		})
	}
}

func updateUserInfo(t *testing.T) {
	testData := []struct {
		Token    string
		Name     string
		Password string
		Avatar   string
		HttpCode int
		Response string
	}{
		{"我是万能钥匙", "哈哈", "", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password为必填字段\"}"},
		{"我是万能钥匙", "", "123456", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name为必填字段\"}"},
		{"我是万能钥匙", "哈哈", "111111", "", http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"哈哈\"},\"msg\":\"success\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := NewMockAuthService(ctrl)
			handler.AuthDomain(service)
			service.EXPECT().Parse(gomock.Any(), gomock.Any()).Return(&authService.ParseResponse{
				Message: "success",
				Data: &authService.User{
					UserId: 1,
					Name:   "张三",
				},
			}, nil)

			byteData, _ := json.Marshal(data)
			reader := bytes.NewReader(byteData)

			resp, err := http.Post("http://127.0.0.1:18282/user/info", "application/json", reader)
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != data.HttpCode {
				t.Errorf("响应HttpCode错误，want:%v, got:%v", data.HttpCode, resp.StatusCode)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != data.Response {
				t.Errorf("响应body错误，want:%v, got:%v", data.Response, string(body))
			}
		})
	}
}
