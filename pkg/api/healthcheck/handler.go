package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PingHandler handler for GET /healthcheck/ping
type PingHandler echo.HandlerFunc

func makePingHandler() PingHandler {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, pingResponse{Ping: "PONG"})
	}
}
