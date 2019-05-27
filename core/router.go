package core

import (
	"logistic-api/core/request"

	"github.com/labstack/echo"
)

func Route(path string, controller IController, method request.Method) {
	if method&request.GET > 0 {
		engine.GET(path, func(c echo.Context) error {
			controller.Init()
			controller.SetContext(c)
			controller.GET()
			return nil
		})
	}
}
