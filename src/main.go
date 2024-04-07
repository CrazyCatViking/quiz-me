package main

import (
	"github.com/CrazyCatViking/quiz-me/src/db"
	"github.com/CrazyCatViking/quiz-me/src/handlers"
	"github.com/CrazyCatViking/quiz-me/src/ioc"
	"github.com/labstack/echo/v4"
)

func main() {
  container := ioc.NewContainer()
 
  ioc.RegisterSingleton[db.DbContext](container, db.Init)

  handlers.RegisterGet("/", handlers.NewUserHandler, container)
  handlers.RegisterGet("/user", handlers.NewUserHandler, container)

  app := echo.New() 

  app.Use(handlers.CreateCustomContextMiddleware(container))
  app.Use(handlers.RouteHandlerMiddleware)

  app.Start(":3000")
}
