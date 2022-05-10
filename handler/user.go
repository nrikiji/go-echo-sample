package handler

import (
	ae "go-echo-starter/error"
	"go-echo-starter/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type usersResponse struct {
	Users []model.User `json:"users"`
}

func (h *Handler) getUsers(c echo.Context) error {
	users, err := h.userStore.All()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, usersResponse{Users: users})
}

type userResponse struct {
	User model.User `json:"user"`
}

func (h *Handler) getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userStore.FindByID(id)

	if err != nil {
		return err
	} else if user == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, userResponse{User: *user})
}

type userRegisterRequest struct {
	Token string `json:"token" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

func (h *Handler) registerUser(c echo.Context) error {
	params := &userRegisterRequest{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err := c.Validate(params); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ae.NewValidationError(err, ae.ValidationMessages{
				"Token": {"required": "Tokenを入力してください"},
				"Name":  {"required": "ユーザー名を入力してください"},
			}),
		)
	}

	uid, err := h.authStore.VerifyIdToken(params.Token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	user, err := h.userStore.FindByUid(uid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if user != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	user = &model.User{
		FirebaseUid: uid,
		Name:        params.Name,
	}

	if err := h.userStore.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, userResponse{User: *user})
}

type userLoginRequest struct {
	Token string `json:"token" validate:"required" ja:"Token"`
}

func (h *Handler) loginUser(c echo.Context) error {
	params := &userLoginRequest{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	uid, err := h.authStore.VerifyIdToken(params.Token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	user, err := h.userStore.FindByUid(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else if user == nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	return c.JSON(http.StatusOK, userResponse{User: *user})

}

type updateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *Handler) updateUser(c echo.Context) error {
	params := &updateUserRequest{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	u := c.Get("user")

	if u == nil {
		return c.JSON(http.StatusForbidden, nil)
	}

	if err := c.Validate(params); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ae.NewValidationError(err, ae.ValidationMessages{
				"Name": {"required": "ユーザー名を入力してください"},
			}),
		)
	}

	user := u.(*model.User)
	user.Name = params.Name

	if err := h.userStore.Update(user); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, userResponse{User: *user})
}
