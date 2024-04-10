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
    cc := c.(*models.CustomContext)

    method := cc.Request().Method
    path := cc.Request().URL.Path

    route := method + path

    fmt.Println(route)
 
    requestContext := handlers.NewRequestContext(cc)
 
    ioc.UseInstance[handlers.RequestContext](cc.Scope, requestContext)

    result, ok := cc.Scope.ResolveRouteHandler(route)

    if !ok {
      return next(cc)
    }

    return result()
  }  
}
