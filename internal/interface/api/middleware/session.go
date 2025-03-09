package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/y2aiskni/spotil/internal/interface/api/helper"
)

var (
	ErrUserUnauthorized = errors.New("user unauthorized")
)

type SessionKey string

var (
	SessionKeyUserID SessionKey = "user_id"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			s, err := helper.GetSession(ctx)
			if err != nil {
				return err
			}

			userID, ok := helper.GetFromSession(s, helper.SessionKeyUserID).(uint64)
			if !ok {
				return ErrUserUnauthorized
			}

			// ToDo: is exist user

			helper.SetToContext(ctx, helper.ContextKeyUserID, userID)
			ctx.Set("user_id", userID)
			return next(ctx)
		}
	}
}
