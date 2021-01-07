package test

import (
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"liaotian/domain/friend/proto"
	"liaotian/domain/friend/repository"
	"net/http"
	"testing"
	"time"
)

func GetFriendList(t *testing.T) {
	testData := []struct {
		UserId int64
		Code   int32
		Msg    string
		Data   string
	}{
		{1, http.StatusOK, "success", "[2,3]"},
		{0, http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)
	for i, data := range testData {
		t.Run("", func(t *testing.T) {

			request := &proto.GetFriendListRequest{
				UserId: data.UserId,
			}
			if i == 0 {
				row := sqlmock.NewRows([]string{"id", "user_id_a", "user_id_b", "created_at", "updated_at"}).
					AddRow(1, data.UserId, 2, time.Now().String(), time.Now().String()).
					AddRow(2, data.UserId, 3, time.Now().String(), time.Now().String())
				repository.Repo.MockDb.ExpectQuery("^SELECT \\* FROM `friends`*").
					WillReturnRows(row)
			}
			resp, err := service.GetFriendList(context.Background(), request)
			if err != nil {
				errData := errors.Parse(err.Error())
				if errData.Code != data.Code {
					t.Errorf("响应Code错误, want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误, want:%v, got:%v", data.Msg, errData.Code)
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

func DeleteFriendInfo(t *testing.T) {
	testData := []struct {
		Id   int64
		Code int32
		Msg  string
		Ok   bool
	}{
		{1, http.StatusOK, "success", true},
		{0, http.StatusBadRequest, "参数错误", false},
	}

	service := proto.NewFriendService("domain.friend.service", client.DefaultClient)
	for i, data := range testData {
		t.Run("", func(t *testing.T) {

			request := &proto.DeleteFriendInfoRequest{
				Id: data.Id,
			}
			if i == 0 {
				repository.Repo.MockDb.ExpectExec("^DELETE FROM `friend`*").
					WithArgs(data.Id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}
			resp, err := service.DeleteFriendInfo(context.Background(), request)
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

			if err = repository.Repo.MockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("sqlmock 执行不符合预期: %v", err)
			}
		})
	}

}
