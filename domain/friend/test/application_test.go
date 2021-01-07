package test

import (
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"liaotian/domain/friend/entity"
	"liaotian/domain/friend/proto"
	"liaotian/domain/friend/repository"
	"net/http"
	"testing"
	"time"
)

func CreateApplicationInfo(t *testing.T) {
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

func GetApplicationInfo(t *testing.T) {
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

func PassApplicationInfo(t *testing.T) {
	testData := []struct{
		Id  int64
		Code int32
		Msg  string
		Ok  bool
	} {
		{1, http.StatusOK, "success", true},
		{1, http.StatusTeapot, "申请单状态错误", false},
		{2, http.StatusNotFound, "申请单不存在", false},
		{0, http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)
	for i, data := range testData {
		t.Run("", func(t *testing.T) {
			request := &proto.PassApplicationInfoRequest{
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
				repository.Repo.MockDb.ExpectExec("^UPDATE `application` SET*").
					WillReturnResult(sqlmock.NewResult(1,1))
				repository.Repo.MockDb.ExpectExec("^INSERT INTO `friend`*").
					WillReturnResult(sqlmock.NewResult(1,1))
			}
			if i == 1 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(data.Id, 1, 2, entity.StatusPass, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(row)
				row = sqlmock.NewRows([]string{"id","application_id", "sender_id", "content", "created_at", "updated_at"}).
					AddRow(1, data.Id, 1, "你好啊", time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WithArgs(data.Id).WillReturnRows(row)
			}
			if i == 2 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.PassApplicationInfo(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if data.Ok != resp.Ok {
					t.Errorf("响应Ok错误, want:%v, got:%v", data.Ok, resp.Ok)
				}
			}

			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
		})
	}
}

func RejectApplicationInfo(t *testing.T) {
	testData := []struct{
		Id  int64
		Code int32
		Msg  string
		Ok  bool
	} {
		{1, http.StatusOK, "success", true},
		{1, http.StatusTeapot, "申请单状态错误", false},
		{2, http.StatusNotFound, "申请单不存在", false},
		{0, http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)
	for i, data := range testData {
		t.Run("", func(t *testing.T) {
			request := &proto.RejectApplicationInfoRequest{
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
				repository.Repo.MockDb.ExpectExec("^UPDATE `application` SET*").
					WillReturnResult(sqlmock.NewResult(1,1))
			}
			if i == 1 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(data.Id, 1, 2, entity.StatusPass, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(row)
				row = sqlmock.NewRows([]string{"id","application_id", "sender_id", "content", "created_at", "updated_at"}).
					AddRow(1, data.Id, 1, "你好啊", time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WithArgs(data.Id).WillReturnRows(row)
			}
			if i == 2 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WithArgs(data.Id).WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.RejectApplicationInfo(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if data.Ok != resp.Ok {
					t.Errorf("响应Ok错误, want:%v, got:%v", data.Ok, resp.Ok)
				}
			}

			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
		})
	}
}

func GetApplicationList(t *testing.T) {
	testData := []struct{
		UserId int64
		Code   int32
		Msg    string
		Data   string
	} {
		{1, http.StatusOK, "success", "[{\"Id\":1,\"SenderId\":1,\"ReceiverId\":2,\"SayList\":[{\"Id\":1,\"SenderId\":1,\"Content\":\"你好啊\"},{\"Id\":2,\"SenderId\":1,\"Content\":\"我是赛利亚\"}]},{\"Id\":2,\"SenderId\":3,\"ReceiverId\":1}]"},
		{2, http.StatusOK, "success", "null"},
		{0, http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)

	for i, data := range testData {
		t.Run("", func(t *testing.T) {
			request := &proto.GetApplicationListRequest{
				UserId: data.UserId,
			}

			if i == 0 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(1, data.UserId, 2, entity.StatusWait, time.Now().String(), time.Now().String()).
					AddRow(2, 3, data.UserId, entity.StatusPass, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WillReturnRows(row)

				row = sqlmock.NewRows([]string{"id","application_id", "sender_id", "content", "created_at", "updated_at"}).
					AddRow(1, 1, 1, "你好啊", time.Now().String(), time.Now().String()).
					AddRow(2, 1, 1, "我是赛利亚", time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WithArgs(1).
					WillReturnRows(row)
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WithArgs(2).
					WillReturnRows(sqlmock.NewRows(nil))
			}
			if i == 1 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.GetApplicationList(context.Background(), request)
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

func CreateApplicationSay(t *testing.T) {
	testData := []struct{
		ApplicationId int64
		SenderId int64
		Content string
		Code   int32
		Msg  string
		Data  string
	} {
		{1, 1, "加我一下吧", http.StatusOK, "success", "{\"Id\":1,\"SenderId\":1,\"Content\":\"加我一下吧\"}"},
		{1, 1,"赶紧加我",http.StatusTeapot, "申请单状态错误", ""},
		{2, 1,"赶紧加我",http.StatusNotFound, "申请单不存在", ""},
		{0, 1,"赶紧加我",http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)

	for i, data := range testData {
		t.Run("", func(t *testing.T) {
			request := &proto.CreateApplicationSayRequest{
				ApplicationId: data.ApplicationId,
				SenderId: data.SenderId,
				Content: data.Content,
			}

			if i == 0 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(1, data.SenderId, 2, entity.StatusWait, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WillReturnRows(row)
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WillReturnRows(sqlmock.NewRows(nil))
				repository.Repo.MockDb.ExpectExec("^INSERT INTO `application_say`*").
					WillReturnResult(sqlmock.NewResult(1,1))
			}
			if i == 1 {
				row := sqlmock.NewRows([]string{"id","sender_id", "receiver_id", "status", "created_at", "updated_at"}).
					AddRow(1, data.SenderId, 2, entity.StatusReject, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WillReturnRows(row)
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `says`*").
					WillReturnRows(sqlmock.NewRows(nil))
			}
			if i == 2 {
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `applications`*").
					WillReturnRows(sqlmock.NewRows(nil))
			}

			resp, err := service.CreateApplicationSay(context.Background(), request)

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
