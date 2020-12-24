package handler

import (
	"liaotian/domain/user/entity"
	"liaotian/domain/user/repository"
)

/**
	领域服务层
 */

type Handler struct {
	UserEntity entity.UserInterface
}

func Init() (handler *Handler) {
	repository.Init()
	handler = new(Handler)
	handler.UserEntity = new(entity.User)
	return
}