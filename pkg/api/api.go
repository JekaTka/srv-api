package api

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Group struct {
	dig.Out

	Group *echo.Group `group:"echo-groups"`
}

type Groups struct {
	dig.In

	Group []*echo.Group `group:"echo-groups"`
}
