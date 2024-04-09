package models

import (
	"github.com/CrazyCatViking/quiz-me/ioc"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
  echo.Context
  Scope *ioc.ContainerScope
}

func NewCustomContext(c echo.Context, scope *ioc.ContainerScope) *CustomContext {
  return &CustomContext{ c, scope }
}
