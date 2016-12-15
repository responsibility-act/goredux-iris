package main

import (
	_ "github.com/adred/goredux/handlers"
	"github.com/adred/goredux/shared"
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func addRoutes(app *iris.Framework) {
	for _, val := range shared.RouteList {
		app.HandleFunc(val.HttpMethod, val.Path, val.HandlerFn)
	}
}

func main() {
	// App config
	config := iris.Configuration{
		IsDevelopment:     true,
		DisablePathEscape: true,
		Gzip:              true,
		Sessions:          iris.SessionsConfiguration{Cookie: "goredux"},
	}
	// Create new app
	app := iris.New(config)
	// Logger
	app.Use(logger.New())
	// Add routes
	addRoutes(app)
	// Serve assets
	app.StaticServe("./public", "/assets")
	// Run app
	app.Listen("localhost:8000")
}
