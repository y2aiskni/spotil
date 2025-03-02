package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/y2aiskni/spotil/internal/interface/api/dto"
)

type HealthHandlerInterface interface {
	handlerInterface
	ping(ctx echo.Context) error
}

type healthHandler struct {
}

func NewHealthHandler() HealthHandlerInterface {
	return healthHandler{}
}

func (h healthHandler) RegisterRoutes(base *echo.Group) {
	health := base.Group("/health")

	health.POST("/ping", h.ping)
}

func (h healthHandler) ping(ctx echo.Context) error {
	res := dto.HealthPingResponse{
		Message: "pong!",
	}

	return ctx.JSON(http.StatusOK, res)
}
