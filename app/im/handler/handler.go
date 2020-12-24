package handler

import (
	"github.com/micro/go-micro/client"
	userService "liaotian/domain/user/proto"
)

var (
	domainUser userService.UserService
)

func Init() {
	domainUser = userService.NewUserService("domain.user.service", client.DefaultClient)
}

