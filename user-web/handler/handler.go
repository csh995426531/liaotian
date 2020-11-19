package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"

	user "liaotian/user-service/proto/user"

	"github.com/micro/go-micro/client"
)

var (
	rpcUser user.UserService
)

func Init() {
	rpcUser = user.NewUserService("user.service.user", client.DefaultClient)
}

func Login(ctx *gin.Context) {
	resultCode := http.StatusOK
	resultData := gin.H{}

	var request user.Request
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Error(err)
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("传参错误：%+v", err),
		}
		return
	}

	res, err := rpcUser.Get(ctx, &request)

	if err != nil {
		log.Error(err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	if res.Code != 200 {
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	if res.User.Password == request.Password && res.User.Name == request.Name {
		resultData = gin.H{
			"message": res.Message,
			"user":    res.User,
		}
	} else {
		resultCode = http.StatusUnauthorized
		resultData = gin.H{
			"message": "用户名密码错误",
		}
	}

	defer ctx.JSON(resultCode, resultData)
}
