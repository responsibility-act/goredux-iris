package handlers

import (
	"github.com/adred/goredux/shared"
	"github.com/kataras/iris/v12"
)

func init() {
	RootPath := "/"
	ihObj := indexHandler{}

	shared.MakeRoute(
		shared.Route{HttpMethod: "GET", Path: RootPath, HandlerFn: ihObj.index},
	)
}

type indexHandler struct{}

func (ih *indexHandler) index(c iris.Context) {
	c.View("index.html", iris.Map{"url": "http://" + c.Host()})
}
