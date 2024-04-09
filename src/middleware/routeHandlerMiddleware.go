package middleware

import (
	"github.com/CrazyCatViking/quiz-me/src/handlers"
	"github.com/CrazyCatViking/quiz-me/src/ioc"
	"github.com/CrazyCatViking/quiz-me/src/models"
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
