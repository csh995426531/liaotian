package handler
//
//import (
//	"fmt"
//	"github.com/micro/cli"
//	"github.com/micro/go-micro/web"
//	"liaotian/middlewares/logger/zap"
//	"liaotian/user-web/handler"
//	"net/http"
//	"testing"
//	"net/http/httptest"
//)
//
//func TestMain(m *testing.M) {
//
//	zap.InitLogger()
//
//	//初始化路由
//	ginRouter := handler.InitRouters()
//
//	// create new web handler
//	service := web.NewService(
//		web.Name("app.im.service"),
//		web.Version("latest"),
//		web.Handler(ginRouter),
//	)
//
//	// 服务初始化
//	if err := service.Init(
//		web.Action(func(c *cli.Context) {
//			handler.Init()
//		}),
//	); err != nil {
//		panic(fmt.Sprintf("服务初始化失败，error: %v", err))
//	}
//
//	// run handler
//	if err := service.Run(); err != nil {
//		panic(fmt.Sprintf("服务启动失败，error: %v", err))
//	}
//	fmt.Println("服务启动成功")
//	m.Run()
//}
//
//func TestRegister(t *testing.T) {
//
//	req, err := http.NewRequest("POST", "/user/register", nil)
//	if err != nil {
//		t.Fatalf("NewRequest error: %v", err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(Register)
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	// Check the response body is what we expect.
//	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//
//}