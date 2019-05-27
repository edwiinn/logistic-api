package core

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var engine *echo.Echo = echo.New()

func Run() {
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	engine.Logger.Fatal(engine.Start(":8080"))
}
