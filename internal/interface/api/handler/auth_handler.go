package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/y2aiskni/spotil/internal/interface/api/helper"
	"github.com/y2aiskni/spotil/internal/interface/api/middleware"
)

type AuthHandlerInterface interface {
	handlerInterface
	signin(ctx echo.Context) error
	info(ctx echo.Context) error
}

type authHandler struct {
}

func NewAuthHandler() AuthHandlerInterface {
	return authHandler{}

}

func (h authHandler) RegisterRoutes(base *echo.Group) {
	auth := base.Group("/auth")

	auth.POST("/signin", h.signin)
	auth.GET("/info", h.info, middleware.Auth())
}

func (h authHandler) signin(ctx echo.Context) error {
	s, err := helper.GetSession(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	helper.SetToSession(s, helper.SessionKeyUserID, uint64(1))

	if err := helper.SaveSession(ctx, s); err != nil {
		return ctx.NoContent(http.StatusServiceUnavailable)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h authHandler) info(ctx echo.Context) error {
	userID, ok := helper.GetFromContext(ctx, helper.ContextKeyUserID).(uint64)
	if !ok {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userID,
	})
}
