package main

import (
	"context"
	"go-echo-starter/db"
	"go-echo-starter/env"
	"go-echo-starter/handler"
	"go-echo-starter/store"
	"go-echo-starter/validator"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())

	api := e.Group("/api")

	opt := option.WithCredentialsFile("firebase_secret_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	as := store.NewAuthStore(client)

	d := db.New(env.NewAppEnv())
	us := store.NewUserStore(d)

	h := handler.NewHandler(*as, *us)
	h.Register(api)
	e.Logger.Fatal(e.Start(":8000"))
}
