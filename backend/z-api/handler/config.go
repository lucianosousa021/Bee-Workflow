package handler

import "zapi/usecase"

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(u *usecase.Usecase) *Handler {
	return &Handler{usecase: u}
}
