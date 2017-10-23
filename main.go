package main

import (
	"app/route"
	"app/handler"

	// "github.com/labstack/echo/middleware"
)

func main() {

	router := route.Init()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())

	router.HTTPErrorHandler = handler.JSONErrorHandler

	router.Start(":8080")
	
}
