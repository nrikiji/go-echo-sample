package handler

import (
	ae "go-echo-starter/error"
	"go-echo-starter/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type postsResponse struct {
	Posts []model.Post `json:"posts"`
}

func (h *Handler) getPosts(c echo.Context) error {
	posts, err := h.userStore.AllPosts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, postsResponse{Posts: posts})
}

type postResponse struct {
	Post model.Post `json:"post"`
}

type updatePostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func (h *Handler) updatePost(c echo.Context) error {
	u := c.Get("user")

	if u == nil {
		return c.JSON(http.StatusForbidden, nil)
	}

	user := u.(*model.User)

	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.userStore.FindPostByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if post == nil {
		return c.JSON(http.StatusNotFound, nil)
	} else if post.UserID != user.ID {
		return c.JSON(http.StatusForbidden, nil)
	}

	params := &updatePostRequest{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err := c.Validate(params); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ae.NewValidationError(err, ae.ValidationMessages{
				"Title": {"required": "タイトルを入力してください"},
				"Body":  {"required": "本文を入力してください"},
			}),
		)
	}

	post.Body = params.Body
	post.Title = params.Title

	if err := h.userStore.UpdatePost(post); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, postResponse{Post: *post})
}
