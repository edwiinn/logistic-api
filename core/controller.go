package core

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	Controller struct {
		Ctx echo.Context
	}

	IController interface {
		SetContext(ctx echo.Context)
		GET()
		Init()
	}
)

func (c *Controller) Init() {
	c.Ctx = nil
}

func (c *Controller) GET() {
	echo.NewHTTPError(http.StatusMethodNotAllowed)
}

func (c *Controller) SetContext(ctx echo.Context) {
	c.Ctx = ctx
}
