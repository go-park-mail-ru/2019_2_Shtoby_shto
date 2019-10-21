package security

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

type HandlerSecurity interface {
	CheckSession(h echo.HandlerFunc) echo.HandlerFunc
	CreateSession(w http.ResponseWriter, userID customType.StringUUID) error
	Logout(ctx echo.Context) error
}

type service struct {
	Sm                 *SessionManager
	NotSecurityRouters map[string]struct{}
	mx                 sync.Mutex
}

func CreateInstance(sm *SessionManager) HandlerSecurity {
	return &service{
		Sm: sm,
		NotSecurityRouters: map[string]struct{}{
			"/users/registration": {},
			"/login":              {},
			"/swagger/index.html": {},
		},
	}
}

func (s *service) Logout(ctx echo.Context) error {
	if err := s.Sm.Delete(ctx.Request().Context()); err != nil {
		return err
	}
	return nil
}

func (s *service) CreateSession(w http.ResponseWriter, userID customType.StringUUID) error {
	sessionId, err := s.Sm.Create(userID)
	if err != nil {
		return err
	}
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId.ID.String(),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	return nil
}

func (s service) checkNotSecurity(route string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()
	_, ok := s.NotSecurityRouters[route]
	return ok
}

func (s *service) CheckSession(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		if s.checkNotSecurity(ctx.Request().RequestURI) {
			return h(ctx)
		}
		cookieSessionID, err := ctx.Cookie("session_id")
		if err == http.ErrNoCookie {
			errors.ErrorHandler(ctx.Response(), "No session_id", http.StatusUnauthorized, err)
			return err
		} else if err != nil {
			errors.ErrorHandler(ctx.Response(), "Error cookie", http.StatusUnauthorized, err)
			return err
		}
		ctx.Set("session_id", cookieSessionID.Value)
		ctx.Logger().Info(ctx.Request().Host, ctx.Request().RequestURI)
		if err := s.Sm.Check(&ctx); err != nil {
			errors.ErrorHandler(ctx.Response(), "Error check session", http.StatusUnauthorized, err)
			return err
		}
		return h(ctx)
	}
}
