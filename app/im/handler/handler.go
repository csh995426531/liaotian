package handler

import (
	"github.com/micro/go-micro/client"
	authService "liaotian/domain/auth/proto"
	friendService "liaotian/domain/friend/proto"
	userService "liaotian/domain/user/proto"
)

var (
	domainUser   userService.UserService
	domainFriend friendService.FriendService
	domainAuth   authService.AuthService
)

func Init() {
	UserDomain(userService.NewUserService("domain.user.service", client.DefaultClient))
	FriendDomain(friendService.NewFriendService("domain.friend.service", client.DefaultClient))
	AuthDomain(authService.NewAuthService("domain.auth.service", client.DefaultClient))
}

// 用户领域服务
func UserDomain(service userService.UserService) {
	domainUser = service
}

// 好友领域服务
func FriendDomain(service friendService.FriendService) {
	domainFriend = service
}

// 认证领域服务
func AuthDomain(service authService.AuthService) {
	domainAuth = service
}
