package shared

import "github.com/kataras/iris/v12"

type Route struct {
	HttpMethod string
	Path       string
	HandlerFn  iris.Handler
}

var RouteList []Route

func MakeRoute(route Route) {
	RouteList = append(RouteList, route)
}
