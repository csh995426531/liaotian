package handler

import "liaotian/user-service/repository"

type Handler struct {
	repo repository.Interface
}

func New(repo repository.Interface) (handler *Handler) {
	handler = new(Handler)
	handler.repo = repo
	return handler
}
