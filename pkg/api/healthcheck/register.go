package healthcheck

import (
	"github.com/JekaTka/srv-api/pkg/api"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func Register(c *dig.Container) {
	c.Provide(setupRoutes)
	c.Provide(makePingHandler)
}

func setupRoutes(router *echo.Echo,
	handlePing PingHandler,
) api.Group {
	g := router.Group("/healthcheck")
	g.GET("/ping", echo.HandlerFunc(handlePing))

	return api.Group{Group: g}
}
