package handler

import (
	"liaotian/domain/message/repository"
)

type Handler struct {
}

func Init() (handler *Handler) {
	handler = new(Handler)
	if err := repository.Init(); err != nil {
		panic(err)
	}
	return
}
