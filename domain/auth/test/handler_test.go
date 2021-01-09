package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"liaotian/domain/auth/handler"
	"liaotian/domain/auth/proto"
	"net/http"

	"github.com/micro/go-micro"
	"liaotian/middlewares/logger/zap"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	zap.InitLogger()

	// 新建服务
	service := micro.NewService(
		micro.Name("domain.auth.service"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*15),
	)

	// 注册服务
	_ = proto.RegisterAuthHandler(service.Server(), new(handler.Handler))

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

func TestGeneratedAndParse(t *testing.T) {

	testData := []struct {
		UserId int64
		Name   string
		Code   int32
		Msg    string
		Data   string
	}{
		{1, "张三", http.StatusCreated, "success", "{\"Data\":{\"UserId\":1,\"Name\":\"张三\"},\"Message\":\"success\"}"},
		{0, "张三", http.StatusBadRequest, "参数错误", ""},
		{1, "", http.StatusBadRequest, "参数错误", ""},
	}

	service := proto.NewAuthService("domain.auth.service", client.DefaultClient)

	for _, data := range testData {
		t.Run("", func(t *testing.T) {

			request := proto.GeneratedRequest{
				UserId: data.UserId,
				Name:   data.Name,
			}
			resp, err := service.Generated(context.Background(), &request)
			if err != nil {
				errData := errors.Parse(err.Error())

				if errData.Code != data.Code {
					t.Errorf("响应Code错误，want:%v, got:%v", data.Code, errData.Code)
				}
				if errData.Detail != data.Msg {
					t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, errData.Detail)
				}
			} else {
				if resp.Data == "" {
					t.Error("响应Data错误，不可为空")
				} else {
					request2 := proto.ParseRequest{
						Token: resp.Data,
					}

					resp2, err2 := service.Parse(context.Background(), &request2)
					if err2 != nil {
						errData2 := errors.Parse(err2.Error())

						if errData2.Code != data.Code {
							t.Errorf("响应Code错误，want:%v, got:%v", data.Code, errData2.Code)
						}
						if errData2.Detail != data.Msg {
							t.Errorf("响应Msg错误，want:%v, got:%v", data.Msg, errData2.Detail)
						}
					} else {
						byteData, _ := json.Marshal(resp2)
						if string(byteData) != data.Data {
							t.Errorf("响应Data错误，want:%v, got:%v", data.Msg, string(byteData))
						}
					}
				}
			}
		})
	}
}
