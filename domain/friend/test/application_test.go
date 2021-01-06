package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"liaotian/domain/friend/entity"
	"liaotian/domain/friend/handler"
	"liaotian/domain/friend/proto"
	"liaotian/domain/friend/repository"
	"liaotian/middlewares/logger/zap"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()

	db, mockDb := repository.NewMockDb()
	repository.Init(db, mockDb)

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.friend.service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
	)

	// 注册服务
	_ = proto.RegisterFriendHandler(service.Server(), handler.Init())

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

func TestCreateApplicationInfo(t *testing.T) {
	testData := []struct {
		SenderId   int64
		ReceiverId int64
		Code       int32
		Msg        string
		Data       string
	}{
		{1, 2, http.StatusOK, "success", "{\"Id\":1,\"SenderId\":1,\"ReceiverId\":2}"},
		{0, 2, http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)

	for i, data := range testData {
		t.Run("", func(t *testing.T) {

			request := proto.CreateApplicationRequest{
				SenderId:   data.SenderId,
				ReceiverId: data.ReceiverId,
			}

			if i == 0 {
				repository.Repo.MockDb.ExpectBegin()
				repository.Repo.MockDb.ExpectExec("^INSERT INTO `application`*").
					WillReturnResult(sqlmock.NewResult(1, 1))
				repository.Repo.MockDb.ExpectCommit()
			}

			resp, err := service.CreateApplicationInfo(context.Background(), &request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误，want:%v，got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if data.Data != "" {
					byteData, _ := json.Marshal(resp.Data)
					if string(byteData) != data.Data {
						t.Errorf("响应Data错误，want:%v, got:%v", data.Data, string(byteData))
					}
				}
			}

			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
		})
	}
}

func TestGetApplicationInfo(t *testing.T) {
	testData := []struct {
		Id   int64
		Code int32
		Msg  string
		Data string
	}{
		{1, http.StatusOK, "success", "{\"Id\":1,\"SenderId\":1,\"ReceiverId\":2,\"SayList\":[{\"Id\":1,\"SenderId\":1,\"Content\":\"你好啊\"}]}"},
		{2, http.StatusNotFound, "申请单不存在", ""},
		{0, http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)

	for i, data := range testData {
		t.Run("", func(t *testing.T) {

			request := proto.GetApplicationRequest{
				Id: data.Id,
			}

			if i == 0 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(data.Id, 1, 2, entity.StatusWait, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(row)
				row = sqlmock.NewRows([]string{"id","application_id", "sender_id", "content", "created_at", "updated_at"}).
					AddRow(1, data.Id, 1, "你好啊", time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WithArgs(data.Id).WillReturnRows(row)
			}
			if i == 1 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.GetApplicationInfo(context.Background(), &request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if data.Data != "" {
					byteData, _ := json.Marshal(resp.Data)
					if string(byteData) != data.Data {
						t.Errorf("响应Data错误, want:%v, got:%v", data.Data, string(byteData))
					}
				}
			}

			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
		})
	}
}

func TestPassApplicationInfo(t *testing.T) {

}

func TestRejectApplicationInfo(t *testing.T) {

}

func TestGetApplicationList(t *testing.T) {

}

func TestCreateApplicationSay(t *testing.T) {

}
