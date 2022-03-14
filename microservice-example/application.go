package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fsosauala/microservice-example/internal/adapters"
	"github.com/fsosauala/microservice-example/internal/core/services"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.String("port", "5000", "the port flag")
	flag.Parse()
	countriesRepository := adapters.NewCountryRepository()
	countriesService := services.NewCountryService(countriesRepository)
	echoHandler := adapters.NewHTTPHandler(countriesService)

	echoHandler.Pre(middleware.RemoveTrailingSlash())
	echoHandler.Use(middleware.CORS())
	echoHandler.Use(middleware.Logger())
	echoHandler.Use(middleware.Recover())
	// Start server
	go func() {
		if err := echoHandler.Start(fmt.Sprintf(":%s", *port)); err != nil {
			echoHandler.Logger.Info("shutting down the server")
		}
	}()
	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := echoHandler.Shutdown(ctx); err != nil {
		echoHandler.Logger.Fatal(err)
	}
}
