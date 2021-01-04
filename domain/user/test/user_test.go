package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/micro/go-micro/client"
	"liaotian/domain/user/handler"
	"liaotian/domain/user/proto"
	"liaotian/domain/user/repository"
	"net/http"

	"github.com/micro/go-micro"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()

	db, mockDb := repository.NewMockDb()
	repository.Init(db, mockDb)

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.user.service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
	)

	// 注册服务
	_ = proto.RegisterUserHandler(service.Server(), handler.Init())

	go func() {
		// 启动服务
		if err := service.Run(); err != nil {
			zap.SugarLogger.Fatalf("服务启动失败，error: %v", err)
		}
	}()

	fmt.Print("服务启动成功")
	time.Sleep(time.Second * 1)
	m.Run()
}

func TestCreateUserInfo(t *testing.T) {

	testData := []struct{
		Account  string
		Name     string
		Password string
		Avatar   string
		Code 	 int64
		Msg      string
		Data     string
	} {
		{"zhangsan", "张三", "123456", "http://baidu.com", http.StatusCreated, "success", "{\"id\":1,\"name\":\"张三\",\"password\":\"123456\",\"avatar\":\"http://baidu.com\"}"},
		{"zhangsan", "张三", "123456", "", http.StatusForbidden, "账户已被注册！", ""},
		{"", "张三", "123456", "http://baidu.com", http.StatusBadRequest, "缺少参数", ""},
		{"zhangsan", "", "123456", "http://baidu.com", http.StatusBadRequest, "缺少参数", ""},
		{"zhangsan", "张三", "", "http://baidu.com", http.StatusBadRequest, "缺少参数", ""},
	}

	service := proto.NewUserService("domain.user.service", client.DefaultClient)

	for i, data := range testData {
		t.Run("", func(t *testing.T) {

			request := proto.Request{
				Account: data.Account,
				Name: data.Name,
				Password: data.Password,
				Avatar: data.Avatar,
			}

			if i == 0 {
				repository.Repo.MockDb.ExpectQuery("SELECT \\* FROM `users`").
					WithArgs(data.Account).
					WillReturnRows(sqlmock.NewRows(nil))
				repository.Repo.MockDb.ExpectBegin()
				repository.Repo.MockDb.ExpectExec("^INSERT INTO `users` (`name`,`account`,`password`,`avatar`,`created_at`,`updated_at`)*").
					WillReturnResult(sqlmock.NewResult(1, 1))
				repository.Repo.MockDb.ExpectCommit()
			}

			if i == 1 {
				row := sqlmock.NewRows([]string{"id", "name", "account", "password", "avatar"}).
					AddRow(1, data.Name, data.Account, data.Password, data.Avatar)
				repository.Repo.MockDb.ExpectQuery("SELECT \\* FROM `users`").
					WithArgs(data.Account).
					WillReturnRows(row)
			}

			resp, err := service.CreateUserInfo(context.Background(), &request)
			if err != nil {
				t.Error(err)
			}
			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期 : %v", err)
			}
			if resp.Code != data.Code {
				t.Errorf("响应Code错误，want:%v, got:%v", data.Code, resp.Code)
			}
			if resp.Message != data.Msg {
				t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, resp.Message)
			}
			if data.Data != "" {
				byteData, _ := json.Marshal(resp.Data)
				if string(byteData) != data.Data {
					t.Errorf("响应Data错误，want:%v, got:%v", data.Data, string(byteData))
				}
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	testData := []struct{
		Account  string
		Name     string
		Id  	 int64
		Code     int64
		Msg      string
		Data     string
	} {
		{"zhangsan", "张三", 1, http.StatusOK, "success", "{\"id\":1,\"name\":\"张三\",\"account\":\"zhangsan\",\"avatar\":\"http://baidu.com\"}"},
		{"lisi", "", 0, http.StatusNotFound, "用户不存在", ""},
		{"", "", 0, http.StatusBadRequest, "缺少参数", ""},
	}

	service := proto.NewUserService("domain.user.service", client.DefaultClient)

	for i, data := range testData{
		t.Run("", func(t *testing.T) {

			request := proto.Request{
				Account: data.Account,
				Name: data.Name,
				Id: data.Id,
			}

			if i == 0 {
				row := sqlmock.NewRows([]string{"id", "name", "account", "password", "avatar"}).
					AddRow(1, data.Name, data.Account, "123456", "http://baidu.com")
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WillReturnRows(row)
			}
			if i == 1 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.GetUserInfo(context.Background(), &request)
			if err != nil {
				t.Error(err)
			}
			if resp.Code != data.Code {
				t.Errorf("响应Code错误，want:%v, got?:%v", data.Code, resp.Code)
			}
			if resp.Message != data.Msg {
				t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, resp.Message)
			}
			if resp.Message != data.Msg {
				t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, resp.Message)
			}
			if data.Data != "" {
				byteData, _ := json.Marshal(resp.Data)
				if string(byteData) != data.Data {
					t.Errorf("响应Data错误，want:%v, got:%v", data.Data, string(byteData))
				}
			}
		})
	}
}

func TestUpdateUserInfo(t *testing.T) {
	testData := []struct{
		Id       int64
		Account  string
		Name     string
		Password string
		Avatar   string
		Code     int64
		Msg      string
		Data 	 string
	} {
		{1, "zhangsan", "张三2", "123123", "http://google.com", http.StatusOK, "success", "{\"id\":1,\"name\":\"张三2\",\"password\":\"123123\",\"avatar\":\"http://google.com\"}"},
		{2, "zhangsan", "张三2", "123123", "http://google.com", http.StatusNotFound, "用户不存在", ""},
		{0, "zhangsan", "张三2", "123123", "http://google.com", http.StatusBadRequest, "缺少参数", ""},
	}

	service := proto.NewUserService("domain.user.service", client.DefaultClient)

	for i, data := range testData{
		t.Run("", func(t *testing.T) {

			if i == 0 {
				row := sqlmock.NewRows([]string{"id","name","account","password","avatar"}).
					AddRow(data.Id, data.Name, data.Account, data.Password, data.Avatar)
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WithArgs(data.Id).
					WillReturnRows(row)
				repository.Repo.MockDb.ExpectBegin()
				repository.Repo.MockDb.ExpectExec("^UPDATE `users` SET*").
					WillReturnResult(sqlmock.NewResult(1, 1))
				repository.Repo.MockDb.ExpectCommit()
			}

			if i == 1 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WithArgs(data.Id).
					WillReturnRows(sqlmock.NewRows(nil))
			}

			request := proto.Request{
				Id: data.Id,
				Name: data.Name,
				Password: data.Password,
				Avatar: data.Avatar,
			}
			resp, err := service.UpdateUserInfo(context.Background(), &request)
			if err != nil {
				t.Error(err)
			}
			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
			if resp.Code != data.Code {
				t.Errorf("响应Code错误，want:%v, got:%v", data.Code, resp.Code)
			}
			if resp.Message != data.Msg {
				t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, resp.Message)
			}
			if data.Data != "" {
				byteData, _ := json.Marshal(resp.Data)
				if string(byteData) != data.Data {
					t.Errorf("响应Data错误，want:%v, got:%v", data.Data, string(byteData))
				}
			}
		})
	}

}

func TestCheckUserPwd(t *testing.T) {
	testData := []struct{
		Id           int64
		Account      string
		Name         string
		Password     string
		Avatar       string
		Code		 int64
		Msg 		 string
		Data 		 string
	}{
		{1, "zhangsan", "张三", "123456","http://baidu.com", http.StatusOK, "success", "{\"id\":1,\"name\":\"张三\",\"password\":\"123456\",\"avatar\":\"http://baidu.com\"}"},
		{1, "zhangsan", "张三", "111111","", http.StatusUnauthorized, "密码错误", ""},
		{2, "lisi", "李四", "123456","http://baidu.com", http.StatusNotFound, "用户不存在", ""},
		{0, "", "张三", "123456","http://baidu.com", http.StatusBadRequest, "缺少参数", ""},
	}

	service := proto.NewUserService("domain.user.service", client.DefaultClient)
	for i, data := range testData{
		t.Run("", func(t *testing.T) {
			request := proto.Request{
				Account: data.Account,
				Password: data.Password,
			}
			if i < 2 {
				row := sqlmock.NewRows([]string{"id","name","account","password","avatar"}).
					AddRow(data.Id, data.Name, data.Account, "123456", data.Avatar)
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WithArgs(data.Account).
					WillReturnRows(row)
			}
			if i == 2 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `users`*").
					WithArgs(data.Account).
					WillReturnRows(sqlmock.NewRows(nil))
			}
			resp, err := service.CheckUserPwd(context.Background(), &request)
			if err != nil {
				t.Error(err)
			}
			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
			if resp.Code != data.Code {
				t.Errorf("响应Code错误，want:%v, got:%v", data.Code, resp.Code)
			}
			if resp.Message != data.Msg {
				t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, resp.Message)
			}
			if data.Data != "" {
				byteData, _ := json.Marshal(resp.Data)
				if string(byteData) != data.Data {
					t.Errorf("响应Data错误，want:%v, got:%v", data.Data, string(byteData))
				}
			}
		})
	}
}



