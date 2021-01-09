package handler

import (
	authService "liaotian/domain/auth/proto"
	friendService "liaotian/domain/friend/proto"
	userService "liaotian/domain/user/proto"
)

var (
	domainUser   userService.UserService
	domainFriend friendService.FriendService
	domainAuth   authService.AuthService
)

func UserDomain(service userService.UserService) {
	domainUser = service
}

func FriendDomain(service friendService.FriendService) {
	domainFriend = service
}

func AuthDomain(service authService.AuthService) {
	domainAuth = service
}
