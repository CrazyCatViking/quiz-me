package handlers

import (
	"github.com/CrazyCatViking/quiz-me/ioc"
)

type RouteHandler interface {
  HandleRoute() error 
}

type RouteManager struct {
  container *ioc.Container
}

func (r *RouteManager) RegisterGet(route string, handler interface{}) {
  r.container.RegisterRouteHandler(handler, "GET" + route)
}

func (r *RouteManager) RegisterPost(route string, handler interface{}) {
  r.container.RegisterRouteHandler(handler, "POST" + route)
}

func NewRouteManager(container *ioc.Container) *RouteManager {
  return &RouteManager{container}
}
