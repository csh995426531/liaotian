package handler

import "liaotian/domain/friend/entity"

type Handler struct {
	ApplicationEntity entity.ApplicationInterface
	ApplicationSayEntity entity.SayInterface
	FriendEntity entity.FriendInterface
}

func Init () (handler *Handler) {
	handler = new(Handler)
	handler.ApplicationEntity = new(entity.Application)
	handler.ApplicationSayEntity = new(entity.Say)
	handler.FriendEntity = new(entity.Friend)
	return
}