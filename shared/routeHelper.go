package shared

import "github.com/kataras/iris"

type Route struct {
	HttpMethod string
	Path string
	HandlerFn iris.HandlerFunc
}

var RouteList []Route

func MakeRoute(route Route) {
	RouteList = append(RouteList, route)
}
