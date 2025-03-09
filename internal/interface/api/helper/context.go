package helper

import "github.com/labstack/echo/v4"

type ContextKey string

var (
	ContextKeyUserID ContextKey = "user_id"
)

func SetToContext(ctx echo.Context, key ContextKey, value any) {
	ctx.Set(string(key), value)
}

func GetFromContext(ctx echo.Context, key ContextKey) any {
	return ctx.Get(string(key))
}
