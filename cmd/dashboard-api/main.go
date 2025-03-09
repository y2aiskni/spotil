package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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

	store, err := redistore.NewRediStore(config.Redis.MaxIdle, "tcp", config.Redis.Address, config.Redis.Username, config.Redis.Password, []byte(config.Session.Secret))
	if err != nil {
		return err
	}
	store.SetKeyPrefix(config.Session.Store.Prefix)
	store.Options = &sessions.Options{
		Path:     config.Cookie.Path,
		Domain:   config.Cookie.Domain,
		Secure:   config.Cookie.Secure,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	store.SetMaxAge(config.Cookie.MaxAge)

	e := echo.New()

	// Middlewares
	e.Use(session.Middleware(store))

	// Handlers
	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler()

	// Routes
	v1 := e.Group("/api/v1")
	healthHandler.RegisterRoutes(v1)
	authHandler.RegisterRoutes(v1)

	if err := e.Start(config.Server.Listen); err != nil {
		log.Fatalln(err)
	}

	return nil
}
