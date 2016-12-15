package handlers

import (
	"github.com/adred/goredux/shared"
	"github.com/kataras/iris"
)

func init() {
	RootPath := "/"
	ihObj := indexHandler{}

	shared.MakeRoute(
		shared.Route{HttpMethod: "GET", Path: RootPath, HandlerFn: ihObj.index},
	)
}

type indexHandler struct{}

func (ih *indexHandler) index(c *iris.Context) {
	c.Render("index.html", iris.Map{"url": "http://" + c.HostString()})
}
