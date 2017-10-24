package main

import (
	"app/route"
	"app/handler"
	"app/log"
	"app/model"
	"app/config"
	"flag"
)

func main() {

	setConfig()

	router := route.Init()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())

	model.Init()

	router.HTTPErrorHandler = handler.JSONErrorHandler

	log.AppLog.Info("サーバーが起動しました")

	router.Start(":8080")
	
}

func setConfig() {
	env := "development"
	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
		env = "production"
	}
	config.SetDB(env)
}
