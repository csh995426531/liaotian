package handler

import (
	"liaotian/domain/user/entity"
)

/**
	领域服务层
 */

type Handler struct {
	UserEntity entity.UserInterface
}

func Init() (handler *Handler) {
	handler = new(Handler)
	handler.UserEntity = new(entity.User)
	return
}