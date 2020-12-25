package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userService "liaotian/domain/user/proto"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	resultCode := http.StatusOK
	resultData := gin.H{}

	var request userService.Request
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("传参错误：%+v", err),
		}
		return
	}

	res, err := domainUser.CheckUserPwd(ctx.Request.Context(), &request)

	if err != nil || res.Code != http.StatusOK {
		log.Printf("domainUser.CheckUserPwd error: %+v", err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	resultData = gin.H{
		"message": res.Message,
		"data":    res.Data,
	}

	defer ctx.JSON(resultCode, resultData)
}

func Register(ctx *gin.Context) {
	resultCode := http.StatusOK
	resultData := gin.H{}

	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message" : fmt.Sprintf("传参错误：%+v", err),
		}
		return
	}

	res, err := domainUser.CreateUserInfo(ctx.Request.Context(), &result)
	if err != nil || res.Code != http.StatusOK {
		log.Printf("domainUser.CreateUserInfo error: %+v", err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message" : res.Message,
		}
		return
	}

	resultData = gin.H{
		"message": res.Message,
		"data": res.Data,
	}

	defer ctx.JSON(resultCode, resultData)
}

func GetUserInfo(ctx *gin.Context) {
	resultCode := http.StatusOK
	resultData := gin.H{}

	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("传参错误：%+v", err),
		}
		return
	}

	res, err := domainUser.GetUserInfo(ctx.Request.Context(), &result)
	if err != nil || res.Code != http.StatusOK {
		log.Printf("domainUser.GetUserInfo error: %+v", err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	resultData = gin.H{
		"message": res.Message,
		"data": res.Data,
	}

	defer ctx.JSON(resultCode, resultData)
}

func UpdateUserInfo(ctx *gin.Context) {

	resultCode := http.StatusOK
	resultData := gin.H{}

	var result userService.Request

	err := ctx.ShouldBindJSON(&result)
	if err != nil {
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("传参错误：%+v", err),
		}
		return
	}

	res, err := domainUser.UpdateUserInfo(ctx.Request.Context(), &result)
	if err != nil || res.Code != http.StatusOK {
		log.Printf("domainUser.UpdateUserInfo error: %+v", err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": res.Message,
		}
		return
	}

	resultData = gin.H{
		"message": res.Message,
		"data": res.Data,
	}

	defer ctx.JSON(resultCode, resultData)
}