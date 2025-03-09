package helper

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var cookieKeySession = "session"

func GetSession(ctx echo.Context) (*sessions.Session, error) {
	s, err := session.Get(cookieKeySession, ctx)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type SessionKey string

var (
	SessionKeyUserID SessionKey = "user_id"
)

func GetFromSession(s *sessions.Session, key SessionKey) any {
	return s.Values[string(key)]
}

func SetToSession(s *sessions.Session, key SessionKey, value any) {
	s.Values[string(key)] = value
}

func SaveSession(ctx echo.Context, s *sessions.Session) error {
	return s.Save(ctx.Request(), ctx.Response())
}
