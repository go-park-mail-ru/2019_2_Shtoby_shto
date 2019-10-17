package security

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/errors"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type HandlerSecurity interface {
	CheckSession(h echo.HandlerFunc) echo.HandlerFunc
	CreateSession(w http.ResponseWriter, userID customType.StringUUID) error
	SecurityResponse(w http.ResponseWriter, status int, respMessage string, err error)
}

type service struct {
	Sm *SessionManager
}

type ResponseSecurity struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func CreateInstance(sm *SessionManager) HandlerSecurity {
	return &service{
		Sm: sm,
	}
}

//func (s *service) LogoutEcho(ctx echo.Context) error {
//	err := s.Sm.Delete(ctx.Request().Context())
//	if err != nil {
//		errors.ErrorHandler(ctx.Response(), "Error delete session", http.StatusInternalServerError, err)
//		return err
//	}
//	ctx.Response().Header().Del("session_id")
//	s.SecurityResponse(ctx.Response(), http.StatusOK, "Logout", err)
//	return nil
//}

func (s *service) CreateSession(w http.ResponseWriter, userID customType.StringUUID) error {
	sessionId, err := s.Sm.Create(userID)
	if err != nil {
		return err
	}
	// TODO:: add token in cookie and expire time for session_id
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId.ID.String(),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	return nil
}

func (s *service) SecurityResponse(w http.ResponseWriter, status int, respMessage string, err error) {
	w.WriteHeader(status)
	b, err := json.Marshal(&ResponseSecurity{
		Message: respMessage,
		Error:   err,
	})
	if _, err := w.Write([]byte(b)); err != nil {
		return
	}
}

func (s *service) CheckSession(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		cookieSessionID, err := ctx.Cookie("session_id")
		if err == http.ErrNoCookie {
			errors.ErrorHandler(ctx.Response(), "No session_id", http.StatusUnauthorized, err)
			return err
		} else if err != nil {
			errors.ErrorHandler(ctx.Response(), "Error cookie", http.StatusUnauthorized, err)
			return err
		}
		ctx.Set("session_id", cookieSessionID.Value)
		if err := s.Sm.Check(&ctx); err != nil {
			errors.ErrorHandler(ctx.Response(), "Error check session", http.StatusUnauthorized, err)
			return err
		}
		return h(ctx)
	}
}
