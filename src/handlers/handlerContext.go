package handlers

import (
	"github.com/CrazyCatViking/quiz-me/src/ioc"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
  echo.Context
  scope *ioc.ContainerScope
}

func CreateCustomContextMiddleware(container *ioc.Container) echo.MiddlewareFunc {
  scope := container.OpenScope()

  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      cc := &CustomContext{ c, scope }
      return next(cc)
    }
  }
}
