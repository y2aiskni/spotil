package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/labstack/echo/v4"
	"github.com/y2aiskni/spotil/internal/interface/api/handler"
)

var version = "0.0.0"

func main() {
	ctx := kong.Parse(&args, &kong.Vars{"version": version})
	switch ctx.Command() {
	case "":
		if err := run(); err != nil {
			log.Fatalln(err)
		}
	default:
		panic(ctx.Command())
	}
}

func run() error {
	config, err := NewConfigFromFile()
	if err != nil {
		return err
	}

	e := echo.New()

	// Handlers
	healthHandler := handler.NewHealthHandler()

	// Routes
	v1 := e.Group("/api/v1")
	healthHandler.RegisterRoutes(v1)

	if err := e.Start(config.Server.Listen); err != nil {
		log.Fatalln(err)
	}

	return nil
}
