package main

import (
	"github.com/piyush2206/go-to-do/app"
	"github.com/piyush2206/go-to-do/server"
)

func main() {
	appCtx := app.Init()

	server.Start(appCtx)
}
