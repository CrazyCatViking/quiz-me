package main

import (
	"github.com/CrazyCatViking/quiz-me/db"
	"github.com/CrazyCatViking/quiz-me/handlers"
	"github.com/CrazyCatViking/quiz-me/ioc"
	"github.com/CrazyCatViking/quiz-me/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
  app := echo.New() 
  container := ioc.NewContainer()
 
  ioc.RegisterSingleton[db.DbContext](container, db.Init)

  routeManager := handlers.NewRouteManager(container) 

  routeManager.RegisterGet("/login", handlers.NewLoginHandler)
  routeManager.RegisterGet("/", handlers.NewUserHandler)

  app.Use(middleware.CreateCustomContextMiddleware(container))
  app.Use(middleware.RouteHandlerMiddleware)

  app.Start(":3000")
}
