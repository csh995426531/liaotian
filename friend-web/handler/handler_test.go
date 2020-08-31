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

func TestAdd(t *testing.T) {
	addTests := []struct{
		OperatorId int64
		BuddyId int64
	} {
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
	}

	for _, test := range addTests {
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)

		param := fmt.Sprintf("{OperatorId: \"%d\", BuddyId: \"%d\"}", test.OperatorId, test.BuddyId)
		ctx.Request,_ = http.NewRequest("POST", "/", bytes.NewBufferString(param))
		Add(ctx)

		if resp.Code != 200 {
			b, _ := ioutil.ReadAll(resp.Body)
			t.Error(resp.Code, string(b))
		}
	}
}

func TestList(t *testing.T) {
	listTests := []struct{
		OperatorId  int64
		Offset      int64
		Limit       int64
	} {
		{1, 0, 2},
		{1, 0, 3},
	}

	for _, test := range listTests {
		resp := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(resp)

		param := fmt.Sprintf("{OperatorId: \"%d\", Offset: \"%d\", Limit: \"%d\"}", test.OperatorId, test.Offset, test.Limit)
		ctx.Request, _ = http.NewRequest("GET", "/", bytes.NewBufferString(param))
		List(ctx)

		if resp.Code != 200 {
			b, _ := ioutil.ReadAll(resp.Body)
			t.Error(resp.Code, string(b))
		}

		//for _, friend := range resp.
		//param = fmt.Sprintf()
		//
		//
	}
}
