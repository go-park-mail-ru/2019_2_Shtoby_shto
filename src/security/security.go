package security

import (
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/errors"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Security interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	CheckSession(h http.HandlerFunc) http.HandlerFunc
}

type service struct {
	Sm   *SessionManager
	User user.UserHandler
}

type LoginResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func CreateInstance(sm *SessionManager, user user.UserHandler) Security {
	return &service{
		Sm:   sm,
		User: user,
	}
}

func (s *service) Logout(w http.ResponseWriter, r *http.Request) {
	ok, session := s.check(r, w)
	if !ok || session.ID == "" {
		errors.ErrorHandler(w, "Error unauthorized", http.StatusUnauthorized, nil)
		return
	}
	if err := s.Sm.Delete(session); err != nil {
		errors.ErrorHandler(w, "Error delete session", http.StatusInternalServerError, err)
		return
	}
}

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(24 * time.Hour)

	user := user.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errors.ErrorHandler(w, "Decode error", http.StatusInternalServerError, err)
		return
	}

	sessionId, err := s.Sm.Create(user)
	if err != nil {
		errors.ErrorHandler(w, "Create error", http.StatusInternalServerError, err)
		return
	}

	// TODO:: add token in cookie and expire time for session_id
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId.ID,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	log.Println("log In")

	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(&LoginResponse{
		Message: "Log In",
		Error:   err,
	})
	w.Write([]byte(b))
	//http.Redirect(w, r, "/books", http.StatusFound)
}

func (s *service) CheckSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, _ := s.check(r, w)
		if !ok {
			s.Login(w, r)
		}
		h.ServeHTTP(w, r)
	})
}

func (s *service) check(r *http.Request, w http.ResponseWriter) (bool, *SessionID) {
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println("No session_id", err)
		s.Login(w, r)
	} else if err != nil {
		log.Println("Error cookie", err)
		errors.ErrorHandler(w, "Error cookie", http.StatusUnauthorized, err)
		return false, nil
	}
	ok, err := s.Sm.Check(&SessionID{ID: cookieSessionID.Value})
	if err != nil {
		log.Println("Error check session", err)
		errors.ErrorHandler(w, "Error check session", http.StatusUnauthorized, err)
		return false, nil
	}
	return ok, &SessionID{ID: cookieSessionID.Value}
}