package main

import (
	"app/route"
	"app/handler"
	// "app/log"
	"app/model"
	"app/config"
	"context"
	// "net/http"
	"os"
	"os/signal"
	"time"
	"fmt"
	"flag"
	"github.com/lestrrat/go-server-starter/listener"
)

func main() {

	setConfig()

	e := route.Init()

	// router.Use(middleware.Logger())
	// router.Use(middleware.Recover())

  listeners, err := listener.ListenAll()
  if err != nil {
      panic(err)
  }
	e.Listener = listeners[0]

	model.Init()
	e.HTTPErrorHandler = handler.JSONErrorHandler

	go func() {
		// if err := e.Start(":8080"); err != nil {
		if err := e.Start(""); err != nil {
			panic(fmt.Sprintf("[Error]: %s", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	
}

func setConfig() {
	env := "development"
	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
		env = "production"
	}
	config.SetEnvironment(env)
}
