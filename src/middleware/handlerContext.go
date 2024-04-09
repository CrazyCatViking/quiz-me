package middleware

import (
	"github.com/CrazyCatViking/quiz-me/src/ioc"
	"github.com/CrazyCatViking/quiz-me/src/models"
	"github.com/labstack/echo/v4"
)

func CreateCustomContextMiddleware(container *ioc.Container) echo.MiddlewareFunc {
  scope := container.OpenScope()

  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      cc := models.NewCustomContext(c, scope)
      return next(cc)
    }
  }
}
