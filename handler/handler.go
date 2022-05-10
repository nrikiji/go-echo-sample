package handler

import "go-echo-starter/store"

type Handler struct {
	authStore store.AuthStore
	dataStore store.DataStore
}

func NewHandler(
	as store.AuthStore,
	ds store.DataStore,
) *Handler {
	return &Handler{
		authStore: as,
		dataStore: ds,
	}
}
