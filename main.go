package main

import (
	"app/route"
	"app/handler"
	"app/log"
	"app/model"
)

func main() {

	router := route.Init()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())

	model.Init()

	router.HTTPErrorHandler = handler.JSONErrorHandler

	log.AppLog.Info("サーバーが起動しました")

	router.Start(":8080")
	
}
