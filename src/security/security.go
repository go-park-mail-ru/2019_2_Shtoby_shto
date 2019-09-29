package security

import (
	. "2019_2_Shtoby_shto/src/custom_type"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/utils"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Security interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Registration(w http.ResponseWriter, r *http.Request)
	CheckSession(h http.HandlerFunc) http.HandlerFunc
	UserSecurity(w http.ResponseWriter, r *http.Request)
}

type service struct {
	Sm   *SessionManager
	User user.UserHandler
}

type ResponseSecurity struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func CreateInstance(sm *SessionManager, user user.UserHandler) Security {
	return &service{
		Sm:   sm,
		User: user,
	}
}

// TODO::replace into handler
func (s *service) UserSecurity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getUserSecurity(w, r)
	case http.MethodPut:
		s.putUserSecurity(w, r)
	}
}

func (s *service) putUserSecurity(w http.ResponseWriter, r *http.Request) {
	user := user.User{}
	cookieId, err := r.Cookie("session_id")
	if err != nil {
		errors.ErrorHandler(w, "Session error", http.StatusUnauthorized, err)
		return
	}

	userId, err := s.Sm.getSession(cookieId.Value)
	if err != nil {
		errors.ErrorHandler(w, "Session error", http.StatusBadRequest, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errors.ErrorHandler(w, "Decode error", http.StatusInternalServerError, err)
		return
	}
	if err := s.User.UpdateUser(user, StringUUID(userId)); err != nil {
		errors.ErrorHandler(w, "Update user error", http.StatusBadRequest, err)
		return
	}
	s.securityResponse(w, http.StatusOK, "Update is success", nil)
}

func (s *service) getUserSecurity(w http.ResponseWriter, r *http.Request) {
	cookieId, err := r.Cookie("session_id")
	if err != nil {
		errors.ErrorHandler(w, "Session error", http.StatusUnauthorized, err)
		return
	}

	userId, err := s.Sm.getSession(cookieId.Value)
	if err != nil {
		errors.ErrorHandler(w, "Session error", http.StatusBadRequest, err)
		return
	}

	user, err := s.User.GetUserById(StringUUID(userId))
	if err != nil {
		errors.ErrorHandler(w, "Update user error", http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(&user)
	if _, err := w.Write([]byte(b)); err != nil {
		return
	}
}

func (s *service) Registration(w http.ResponseWriter, r *http.Request) {
	user := user.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errors.ErrorHandler(w, "Decode error", http.StatusInternalServerError, err)
		return
	}
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		errors.ErrorHandler(w, "Error create UUID", http.StatusInternalServerError, err)
		return
	}
	user.ID = StringUUID(id.String())
	if err := s.User.PutUser(user); err != nil {
		errors.ErrorHandler(w, "User not valid", http.StatusBadRequest, err)
		return
	}
	if err := s.createSession(w, user); err != nil {
		errors.ErrorHandler(w, "Create session error", http.StatusInternalServerError, err)
		return
	}
	s.securityResponse(w, http.StatusOK, "Registration is success", err)

}
func (s *service) Logout(w http.ResponseWriter, r *http.Request) {
	ok, session := s.check(r, w)
	if !ok || session.ID == "" {
		errors.ErrorHandler(w, "Error unauthorized", http.StatusUnauthorized, nil)
		return
	}
	err := s.Sm.Delete(session)
	if err != nil {
		errors.ErrorHandler(w, "Error delete session", http.StatusInternalServerError, err)
		return
	}
	s.securityResponse(w, http.StatusOK, "Logout", err)
}

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	curUser := user.User{}

	if err := json.NewDecoder(r.Body).Decode(&curUser); err != nil {
		errors.ErrorHandler(w, "Decode error", http.StatusInternalServerError, err)
		return
	}

	user, err := s.User.GetUserByLogin(curUser.Login)
	if err != nil {
		errors.ErrorHandler(w, "Please, reg yourself", http.StatusUnauthorized, err)
		return
	}

	if strings.Compare(user.Password, curUser.Password) != 0 {
		errors.ErrorHandler(w, "Ne tot password )0))", http.StatusBadRequest, err)
		return
	}

	if err := s.createSession(w, user); err != nil {
		errors.ErrorHandler(w, "Create session error", http.StatusInternalServerError, err)
		return
	}

	s.securityResponse(w, http.StatusOK, "Login", err)
}

func (s *service) createSession(w http.ResponseWriter, user user.User) error {
	expiration := time.Now().Add(24 * time.Hour)
	sessionId, err := s.Sm.Create(user)
	if err != nil {
		return err
	}
	// TODO:: add token in cookie and expire time for session_id
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId.ID.String(),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	return nil
}

func (s *service) securityResponse(w http.ResponseWriter, status int, respMessage string, err error) {
	w.WriteHeader(status)
	b, err := json.Marshal(&ResponseSecurity{
		Status:  status,
		Message: respMessage,
		Error:   err,
	})
	if _, err := w.Write([]byte(b)); err != nil {
		return
	}
}

func (s *service) CheckSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok, _ := s.check(r, w); !ok {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (s *service) check(r *http.Request, w http.ResponseWriter) (bool, *SessionID) {
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		errors.ErrorHandler(w, "No session_id", http.StatusUnauthorized, err)
		return false, nil
	} else if err != nil {
		errors.ErrorHandler(w, "Error cookie", http.StatusUnauthorized, err)
		return false, nil
	}
	ok, err := s.Sm.Check(&SessionID{ID: StringUUID(cookieSessionID.Value)})
	if err != nil {
		errors.ErrorHandler(w, "Error check session", http.StatusUnauthorized, err)
		return false, nil
	}
	return ok, &SessionID{ID: StringUUID(cookieSessionID.Value)}
}
