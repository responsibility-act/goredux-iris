package handlers

import (
	"github.com/adred/goredux/shared"
	"github.com/kataras/iris"
)

func init() {
	RootPath := "/widgets"
	whObj := widgetHandler{}

	shared.MakeRoute(
		shared.Route{HttpMethod: "GET", Path: RootPath, HandlerFn: whObj.index},
	)
}

type widgetHandler struct{}

func (wh *widgetHandler) index(c *iris.Context) {
	c.Render("index.html", iris.Map{"url": "http://" + c.HostString()})
}
