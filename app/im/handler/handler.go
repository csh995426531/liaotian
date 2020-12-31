package handler

import (
	userService "liaotian/domain/user/proto"
)

var (
	domainUser userService.UserService
)

func Init(service userService.UserService) {
	domainUser = service
}

