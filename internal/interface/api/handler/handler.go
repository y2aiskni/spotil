package handler

import (
	"github.com/labstack/echo/v4"
)

type handlerInterface interface {
	RegisterRoutes(g *echo.Group)
}
