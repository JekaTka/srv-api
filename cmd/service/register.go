package main

import (
	"github.com/JekaTka/srv-api/config"
	"github.com/JekaTka/srv-api/db"
	"github.com/JekaTka/srv-api/pkg/api"
	"github.com/JekaTka/srv-api/pkg/api/healthcheck"
	"github.com/JekaTka/srv-api/pkg/services/password"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"net/http"
	"context"
	"fmt"
)

type serverParams struct {
	dig.In
	Router *echo.Echo
	Groups api.Groups // api.Groups is here to only let them created. echo.Group will automatically inject it into echo router
}
type serverFactory func(params serverParams) *http.Server

func makeServerFactory(ctx context.Context, cfg *config.Config) serverFactory {
	return func(params serverParams) *http.Server {
		return &http.Server{
			Addr:         fmt.Sprintf("[::]:%d", cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			Handler:      params.Router,
		}
	}
}

func bootstrapService(ctx context.Context, cfg *config.Config) func(*dig.Container) {
	return func(c *dig.Container) {
		c.Provide(echo.New)
		c.Provide(makeServerFactory(ctx, cfg))

		// config
		c.Provide(func() *config.Config { return cfg })
		c.Provide(func() *config.DB { return cfg.DB })

		// connect to db
		db.Register(c)

		// services
		password.Register(c)

		// api
		healthcheck.Register(c)
	}
}
