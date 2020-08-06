package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JekaTka/cryptohex-api/pkg/api"
	"github.com/JekaTka/cryptohex-api/pkg/api/healthcheck"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type serverParams struct {
	dig.In
	Router *echo.Echo
	Groups api.Groups // api.Groups is here to only let them created. echo.Group will automatically inject it into echo router
}
type serverFactory func(params serverParams) *http.Server

func makeServerFactory(ctx context.Context /*, cfg *config.Config*/) serverFactory {
	log.Println("test: ")
	return func(params serverParams) *http.Server {
		log.Println("Params: ", params.Router.Routes()[0])
		return &http.Server{
			Addr:         fmt.Sprintf("[::]:%d", 1323),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      params.Router,
		}
	}
}

func bootstrapService(ctx context.Context /*, cfg *config.Config*/) func(*dig.Container) {
	return func(c *dig.Container) {
		c.Provide(echo.New)
		c.Provide(makeServerFactory(ctx))

		// api
		healthcheck.Register(c)
	}
}

func main() {
	ctx := context.Background()
	di := dig.New()

	bootstrapService(ctx)(di)
	if err := di.Invoke(func(httpServer *http.Server) error {
		log.Println("Starting service")

		return httpServer.ListenAndServe()
	}); err != nil {
		log.Fatal("Failed to start server")
		panic(err)
	}
	// di.Invoke(func(e *echo.Echo) {
	// 	e.Logger.Fatal(e.Start(":1323"))
	// })
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
