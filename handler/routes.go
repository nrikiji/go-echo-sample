package handler

import (
	"go-echo-starter/middleware"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(api *echo.Group) {
	auth := middleware.AuthMiddleware(h.authStore, h.userStore)

	users := api.Group("/users")
	users.GET("", h.getUsers)
	users.GET("/:id", h.getUser)
	users.POST("/login", h.loginUser)
	users.POST("/register", h.registerUser)
	users.PUT("", h.updateUser, auth)

	api.GET("/posts", h.getPosts)
	api.PATCH("/posts/:id", h.updatePost, auth)
}
