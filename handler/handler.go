package handler

import "go-echo-starter/store"

type Handler struct {
	authStore store.AuthStore
	userStore store.UserStore
}

func NewHandler(
	as store.AuthStore,
	us store.UserStore,
) *Handler {
	return &Handler{
		authStore: as,
		userStore: us,
	}
}
