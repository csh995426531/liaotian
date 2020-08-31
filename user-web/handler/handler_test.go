package handler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestLogin(t *testing.T) {

	loginTests := []struct {
		Name     string
		Password string
	}{
		{"张三", "aa123123123"},
		{"张四", "aa123123123"},
	}

	for _, test := range loginTests {
		t.Run(test.Name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(resp)
			param := fmt.Sprintf("{name: \"%s\", password: \"%s\"}", test.Name, test.Password)
			ctx.Request,_ = http.NewRequest("POST", "/", bytes.NewBufferString(param))
			Login(ctx)

			if resp.Code != 200 {
				b, _ := ioutil.ReadAll(resp.Body)
				t.Error(resp.Code, string(b))
			}
		})
	}



}
