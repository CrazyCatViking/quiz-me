package middleware

import (
	"fmt"

	"github.com/CrazyCatViking/quiz-me/handler"
	"github.com/CrazyCatViking/quiz-me/ioc"
	"github.com/CrazyCatViking/quiz-me/model"
	"github.com/labstack/echo/v4"
)

func RouteHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    cc := c.(*model.CustomContext)

    method := cc.Request().Method
    path := cc.Request().URL.Path

    route := method + path

    fmt.Println(route)
 
    requestContext := handler.NewRequestContext(cc)
 
    ioc.UseInstance[handler.RequestContext](cc.Scope, requestContext)

    result, ok := cc.Scope.ResolveRouteHandler(route)

    if !ok {
      return next(cc)
    }

    return result()
  }  
}
