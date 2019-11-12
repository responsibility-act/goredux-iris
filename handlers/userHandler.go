package handlers

import (
	"github.com/adred/goredux/shared"
	"github.com/kataras/iris/v12"
)

func init() {
	RootPath := "/users"
	uhObj := userHandler{}

	shared.MakeRoute(
		shared.Route{HttpMethod: "GET", Path: RootPath, HandlerFn: uhObj.index},
	)
	shared.MakeRoute(
		shared.Route{HttpMethod: "GET", Path: RootPath + "/:userId", HandlerFn: uhObj.detail},
	)
}

type userHandler struct{}

func (uh *userHandler) index(c iris.Context) {
	c.View("index.html", iris.Map{"url": "http://" + c.Host()})
}

func (uh *userHandler) detail(c iris.Context) {
	c.View("index.html", iris.Map{"url": "http://" + c.Host()})
}
