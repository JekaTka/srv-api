package main

import (
	"context"
	"github.com/JekaTka/srv-api/config"
	"go.uber.org/dig"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	di := dig.New()
	cfg := config.Load()

	bootstrapService(ctx, cfg)(di)
	if err := di.Invoke(func(httpServer *http.Server) error {
		log.Println("Starting service")

		return httpServer.ListenAndServe()
	}); err != nil {
		log.Fatal("Failed to start server")
		panic(err)
	}
}
