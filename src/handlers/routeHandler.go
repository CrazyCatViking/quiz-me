package handlers

import (
	"github.com/CrazyCatViking/quiz-me/src/ioc"
	"github.com/labstack/echo/v4"
)

type RouteHandler interface {
  HandleRoute(context CustomContext) error 
}

func RouteHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    cc := c.(*CustomContext)

    path := cc.Request().URL.Path

    result, ok := ioc.ResolveRequestHandler[RouteHandler](cc.scope, path)

    if !ok {
      return next(cc)
    }

    return (*result).HandleRoute(*cc)
  }  
}

func RegisterGet(route string, handler interface{}, container *ioc.Container) {
  ioc.RegisterRouteHandler[RouteHandler](container, handler, route)
}
