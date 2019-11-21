package security

import (
	"2019_2_Shtoby_shto/src/customType"
	errLib "2019_2_Shtoby_shto/src/errors"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

type HandlerSecurity interface {
	CheckSession(h echo.HandlerFunc) echo.HandlerFunc
	CreateSession(ctx *echo.Context, userID customType.StringUUID) error
	DeleteSession(ctx echo.Context) error
}

type service struct {
	Sm                *SessionManager
	noSecurityRouters map[string]struct{}
	mx                sync.Mutex
}

func CreateInstance(sm *SessionManager) HandlerSecurity {
	return &service{
		Sm: sm,
		noSecurityRouters: map[string]struct{}{
			"/users/registration": {},
			"/login":              {},
			"/swagger/index.html": {},
		},
	}
}

func (s *service) DeleteSession(ctx echo.Context) error {
	sessionID, ok := ctx.Get("session_id").(string)
	if !ok {
		return errors.New("Session ID is not valid ")
	}
	if err := s.Sm.Delete(sessionID); err != nil {
		return err
	}
	ctx.Set("session_id", "")
	return nil
}

func (s *service) CreateSession(ctx *echo.Context, userID customType.StringUUID) error {
	session, err := s.Sm.Create(userID)
	if err != nil {
		return err
	}
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   session.ID.String(),
		Expires: expiration,
	}
	(*ctx).Response().Header().Add(echo.HeaderXCSRFToken, session.CsrfToken)
	http.SetCookie((*ctx).Response(), &cookie)
	return nil
}

func (s service) checkNotSecurity(route string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	_, ok := s.noSecurityRouters[route]
	return ok
}

func (s *service) CheckSession(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if s.checkNotSecurity(ctx.Request().RequestURI) {
			ctx.Set("not_security", "done")
			return h(ctx)
		}
		cookieSessionID, err := ctx.Cookie("session_id")
		if err == http.ErrNoCookie {
			errLib.ErrorHandler(ctx.Response(), "No session_id", http.StatusUnauthorized, err)
			return err
		} else if err != nil {
			errLib.ErrorHandler(ctx.Response(), "Error cookie", http.StatusUnauthorized, err)
			return err
		}
		ctx.Logger().Info(ctx.Request().Host, ctx.Request().RequestURI)
		session, err := s.Sm.Check(cookieSessionID.Value)
		if err != nil {
			errLib.ErrorHandler(ctx.Response(), "Error check session", http.StatusUnauthorized, err)
			return err
		}
		ctx.Set("session_id", cookieSessionID.Value)
		ctx.Set("user_id", session.UserID)
		ctx.Set("csrf_token", session.CsrfToken)
		return h(ctx)
	}
}
