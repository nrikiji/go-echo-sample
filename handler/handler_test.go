package handler

import (
	"context"
	"errors"
	"go-echo-starter/db"
	"go-echo-starter/env"
	"go-echo-starter/store"
	"go-echo-starter/validator"
	"os"
	"path/filepath"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	e *echo.Echo
	h *Handler
)

func init() {
	apath, _ := filepath.Abs("../")
	os.Chdir(apath)
}

type fakeAuthClient struct {
}

func (f *fakeAuthClient) VerifyIDToken(context context.Context, token string) (*auth.Token, error) {
	var uid string
	if token == "ValidToken" {
		uid = "ValidUID"
		return &auth.Token{UID: uid}, nil
	} else if token == "ValidToken1" {
		uid = "ValidUID1"
		return &auth.Token{UID: uid}, nil
	} else {
		return nil, errors.New("Invalid Token")
	}
}

func setup() {
	e = echo.New()
	e.Validator = validator.NewValidator()
	e.Logger.SetLevel(log.DEBUG)
	t := env.NewTestEnv()

	err := db.NewFixtures(t).Load()

	if err != nil {
		log.Fatal(err)
	}

	client := &fakeAuthClient{}
	as := store.NewAuthStore(client)

	d := db.New(t)
	us := store.NewUserStore(d)

	h = NewHandler(*as, *us)
}
