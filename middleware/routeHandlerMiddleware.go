package middleware

import (
	"github.com/CrazyCatViking/quiz-me/handlers"
	"github.com/CrazyCatViking/quiz-me/ioc"
	"github.com/CrazyCatViking/quiz-me/models"
	"github.com/labstack/echo/v4"
)

func RouteHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    cc := c.(*models.CustomContext)

    path := cc.Request().URL.Path
 
    requestContext := handlers.NewRequestContext(cc)
 
    ioc.UseInstance[handlers.RequestContext](cc.Scope, requestContext)

    result, ok := ioc.ResolveRequestHandler[handlers.RouteHandler](cc.Scope, path)

    if !ok {
      return next(cc)
    }

    return (*result).HandleRoute()
  }  
}
