package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	client "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"io/ioutil"
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
		Id: 1,
		Name: in.Name,
		Account: in.Account,
		Password: in.Password,
		Avatar: in.Avatar,
	}

	return out, nil
}

func (c *testService) GetUserInfo(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)
	if in.Account == "" || in.Name == "" || in.Id == 0{
		out.Code = http.StatusBadRequest
		out.Message = "缺少参数！"
		out.Data = nil
		return out, nil
	}

	out.Code = http.StatusOK
	out.Message = "success"
	out.Data = nil

	if in.Id == 1 || in.Account == "zhangsan" || in.Name == "张三" {
		out.Data = &userService.User{
			Id: 1,
			Name: "张三",
			Account: "zhangsan",
			Avatar: "http://www.baicu.com",
		}
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
		Id: 1,
		Name: "张三",
		Account: "zhangsan",
		Avatar: "http://www.baidu.com",
	}
	return out, nil
}
func (c *testService) CheckUserPwd(ctx context.Context, in *userService.Request, opts ...client.CallOption) (*userService.Response, error) {
	out := new(userService.Response)
	if in.Id != 1 {
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
		Id: 1,
		Name: "张三",
		Account: "zhangsan",
		Avatar: "http://www.baidu.com",
	}

	return out, nil
}

type A struct {
	Name string
}
type B struct {
	Name string
}

func TestMain(m *testing.M) {

	zap.InitLogger()
	//translate.Init()

	//初始化路由
	ginRouter := InitRouters()

	// create new web handler
	service := web.NewService(
		web.Name("app.im.service"),
		web.Version("latest"),
		web.Handler(ginRouter),
		web.Address(":18282"),
	)
	Init(new(testService))

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

	testData := []struct{
		Account  string
		Password string
		Name 	 string
		Avatar 	 string
		HttpCode int
		Response string
	} {
		{"zhangsan", "123456", "张三", "http://baidu.com", http.StatusInternalServerError, "{\"data\":null,\"msg\":\"账户已被注册！\"}"},
		{"", "123456", "李四", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Account为必填字段\"}"},
		{"lisi", "", "李四", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Password为必填字段\"}"},
		{"lisi", "123456", "", "http://baidu.com", http.StatusBadRequest, "{\"data\":null,\"msg\":\"Name为必填字段\"}"},
		{"lisi", "123456", "李四", "", http.StatusCreated,"{\"data\":{\"id\":1,\"name\":\"李四\",\"account\":\"lisi\",\"password\":\"123456\"},\"msg\":\"成功\"}"},
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

//func TestLogin(t *testing.T) {
//	testData := []struct{
//		Account string
//		password string
//	}{
//		{"aaa", "123456"},
//	}
//}