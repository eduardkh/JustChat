package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/justchatapp/justchat/handlers"
	"github.com/justchatapp/justchat/services"
)

func main() {
	services.InitServices()
	e := echo.New()

	e.Renderer = handlers.NewTemplateRenderer()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
