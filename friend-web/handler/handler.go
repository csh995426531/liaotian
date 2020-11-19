package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	proto "liaotian/friend-service/proto/friend"
	protoUser "liaotian/user-service/proto/user"
	"net/http"
)

var (
	rpcFriend proto.FriendService
	rpcUser   protoUser.UserService
)

func Init() {
	rpcFriend = proto.NewFriendService("friend.service.friend", client.DefaultClient)
	rpcUser = protoUser.NewUserService("user.service.user", client.DefaultClient)
}

type UserInfo struct {
	id   int64  `json:"id"`
	name string `json:"name"`
}

func Add(ctx *gin.Context) {
	resultCode := http.StatusOK
	resultData := gin.H{}

	var request proto.AddRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Error(err)
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("参数错误：%+v", err),
		}
		return
	}

	res, err := rpcFriend.Add(ctx, &request)
	if err != nil {
		log.Error(err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": fmt.Sprintf("添加失败：%+v", err),
		}
		return
	}

	if res.Code == http.StatusOK {
		resultData = gin.H{
			"message": fmt.Sprint("SUCCESS"),
		}
		return
	}

	defer ctx.JSON(resultCode, resultData)
}

func Del(ctx *gin.Context) {

	resultCode := http.StatusOK
	resultData := gin.H{}

	var request proto.Request

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Error(err)
		resultCode = http.StatusBadRequest
		resultData = gin.H{
			"message": fmt.Sprintf("参数错误：%+v", err),
		}
		return
	}

	res, err := rpcFriend.Del(ctx, &request)
	if err != nil {
		log.Error(err)
		resultCode = http.StatusInternalServerError
		resultData = gin.H{
			"message": fmt.Sprintf("删除失败：%+v", err),
		}
		return
	}

	if res.Code == http.StatusOK {
		resultData = gin.H{
			"message": "SUCCESS",
		}
	}

	defer ctx.JSON(resultCode, resultData)
}

func List(ctx *gin.Context) {

	httpCode := http.StatusOK
	httpData := gin.H{}

	var request proto.ListRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Error(err)
		httpCode = http.StatusBadRequest
		httpData = gin.H{
			"message": fmt.Sprintf("参数错误：%+v", err),
		}
		return
	}

	res, err := rpcFriend.List(ctx, &request)
	if err != nil {
		log.Error(err)
		httpCode = http.StatusInternalServerError
		httpData = gin.H{
			"message": fmt.Sprintf("查询列表失败：%+v", err),
		}
		return
	}

	var List []*UserInfo

	for _, friend := range res.List {

		userRes, err := rpcUser.Get(ctx, &protoUser.Request{Id: friend.UserId})
		if err != nil {
			log.Error(err)
		}
		if userRes == nil || userRes.User == nil {
			log.Errorf("user信息异常：%+v", userRes)
		} else {

			List = append(List, &UserInfo{
				id:   friend.Id,
				name: userRes.User.Name,
			})
		}
	}

	httpData = gin.H{
		"message": "SUCCESS",
		"data":    List,
	}

	defer ctx.JSON(httpCode, httpData)
}

func Info(ctx *gin.Context) {

	httpCode := http.StatusOK
	httpData := gin.H{}

	var request proto.Request
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Error(err)
		httpCode = http.StatusBadRequest
		httpData = gin.H{
			"message": fmt.Errorf("参数错误：%+v", err),
		}
		return
	}

	res, err := rpcFriend.Get(ctx, &request)
	if err != nil {
		log.Error(err)
		httpCode = http.StatusInternalServerError
		httpData = gin.H{
			"message": fmt.Sprintf("查询friend失败：%+v", err),
		}
		return
	}
	if res == nil || res.Friend == nil {
		log.Error(err)
		httpCode = http.StatusInternalServerError
		httpData = gin.H{
			"message": fmt.Sprintf("friend信息异常：%+v", err),
		}
		return
	} else {

		resUser, err := rpcUser.Get(ctx, &protoUser.Request{Id: res.Friend.UserId})
		if err != nil {
			log.Error(err)
			httpCode = http.StatusInternalServerError
			httpData = gin.H{
				"message": fmt.Sprintf("user信息异常：%+v", err),
			}
		}
		httpData = gin.H{
			"message": "SUCCESS",
			"data": &UserInfo{
				id:   resUser.User.Id,
				name: resUser.User.Name,
			},
		}
	}

	defer ctx.JSON(httpCode, httpData)
}
