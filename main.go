package main

import (
	_ "github.com/adred/goredux/handlers"
	"github.com/adred/goredux/shared"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/sessions"
)

func addRoutes(app *iris.Application) {
	for _, val := range shared.RouteList {
		app.Handle(val.HttpMethod, val.Path, val.HandlerFn)
	}
}

func main() {
	// Create new app
	app := iris.New()
	sessManager := sessions.New(sessions.Config{
		Cookie: "goredux",
	})
	app.Use(sessManager.Handler())
	// Logger
	app.Use(logger.New())
	// Add routes
	addRoutes(app)
	// Serve assets
	app.HandleDir("./public", "/assets")
	// Run app
	app.Run(iris.Addr("localhost:8000"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
