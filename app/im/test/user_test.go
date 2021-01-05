package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	client "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"io/ioutil"
	"liaotian/app/im/handler"
	userService "liaotian/domain/user/proto"
	"liaotian/middlewares/logger/zap"
	"net/http"
	"testing"
	"time"
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
		Id:       1,
		Name:     in.Name,
		Account:  in.Account,
		Password: in.Password,
		Avatar:   in.Avatar,
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
			Avatar:  "http://www.baicu.com",
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
	out.Message = "成功"
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
	out.Message = "成功"
	out.Data = &userService.User{
		Id:      1,
		Name:    "张三",
		Account: "zhangsan",
		Avatar:  "http://www.baidu.com",
	}

	return out, nil
}

func TestMain(m *testing.M) {

	zap.InitLogger()
	//translate.Init()

	//初始化路由
	ginRouter := handler.InitRouters()

	// create new web handler
	service := web.NewService(
		web.Name("app.im.service"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(":18282"),
	)
	handler.Init(new(testService))

	// run handler
	go func() {
		if err := service.Run(); err != nil {
			panic(fmt.Sprintf("服务启动失败，error: %v", err))
		}
	}()

	fmt.Println("服务启动成功")
	time.Sleep(time.Second * 1)
	m.Run()
}

func TestRegister(t *testing.T) {

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
		{"lisi", "123456", "李四", "", http.StatusCreated, "{\"data\":{\"id\":1,\"name\":\"李四\",\"account\":\"lisi\",\"password\":\"123456\"},\"msg\":\"成功\"}"},
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

func TestLogin(t *testing.T) {
	testData := []struct {
		Account  string
		Password string
		HttpCode int
		Response string
	}{
		{"zhangsan", "123456", http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"张三\",\"account\":\"zhangsan\",\"avatar\":\"http://www.baidu.com\"},\"msg\":\"成功\"}"},
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

func TestGetUserInfo(t *testing.T) {
	testData := []struct {
		Id       int64
		HttpCode int
		Response string
	}{
		{1, http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"张三\",\"account\":\"zhangsan\",\"avatar\":\"http://www.baicu.com\"},\"msg\":\"成功\"}"},
		{2, http.StatusNotFound, "{\"data\":null,\"msg\":\"用户不存在\"}"},
		{0, http.StatusBadRequest, "{\"data\":null,\"msg\":\"Id为必填字段\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			url := fmt.Sprintf("http://127.0.0.1:18282/user/info?Id=%v", data.Id)
			resp, err := http.Get(url)
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

func TestUpdateUserInfo(t *testing.T) {
	testData := []struct {
		Id       int64
		Name     string
		Password string
		Avatar   string
		HttpCode int
		Response string
	}{
		{1, "哈哈", "", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password为必填字段\"}"},
		{1, "", "123456", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name为必填字段\"}"},
		{1, "哈哈", "111111", "", http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"哈哈\"},\"msg\":\"成功\"}"},
		{0, "哈哈", "111111", "", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Id为必填字段\"}"},
		{2, "哈哈", "111111", "http://baidu.com", http.StatusNotFound, "{\"data\":null,\"msg\":\"用户不存在\"}"},
		{1, "哈哈", "111111", "http://baidu.com", http.StatusOK, "{\"data\":{\"id\":1,\"name\":\"哈哈\",\"avatar\":\"http://baidu.com\"},\"msg\":\"成功\"}"},
		{1, "aaaaaaaaaaaaaaaaaaaaa", "111111", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name长度不能超过20个字符\"}"},
		{1, "哈哈", "11111111111111111111111", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password长度不能超过20个字符\"}"},
	}

	for _, data := range testData {
		t.Run("", func(t *testing.T) {
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
				t.Errorf("相应body错误，want:%v, got:%v", data.Response, string(body))
			}
		})
	}
}
