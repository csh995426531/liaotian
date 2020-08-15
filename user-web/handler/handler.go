package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/v2/logger"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/client"
	user "liaotian/user-service/proto/user"
)

var (
	rpcUser	user.UserService
)

func Init () {
	rpcUser = user.NewUserService("user.service.user", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		logger.Error("非法请求")
		http.Error(w, "非法请求", 500)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	data := user.Request{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := rpcUser.Get(context.TODO(), &data)

	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if res.Code != 200 {
		http.Error(w, res.Message, 500)
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
		"user": &res.User,
		"rpc_data": &res,
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
