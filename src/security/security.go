package security

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/utils"
	"bufio"
	"context"
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
	ImageSecurity(w http.ResponseWriter, r *http.Request)
}

type service struct {
	Sm    *SessionManager
	User  user.HandlerUserService
	Photo photo.HandlerPhotoService
}

type ResponseSecurity struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func CreateInstance(sm *SessionManager, user user.HandlerUserService, p photo.HandlerPhotoService) Security {
	return &service{
		Sm:    sm,
		User:  user,
		Photo: p,
	}
}

// TODO::replace into handler
func (s *service) ImageSecurity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rr := bufio.NewReader(r.Body)
		photoID, err := s.Photo.DownloadPhoto(rr)
		if err != nil {
			errors.ErrorHandler(w, "download fail", http.StatusInternalServerError, err)
			return
		}

		// TODO:: add _context with session and user values
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
			errors.ErrorHandler(w, "GetUserById error", http.StatusInternalServerError, err)
			return
		}
		user.PhotoID = &photoID
		if err := s.User.UpdateUser(user, StringUUID(userId)); err != nil {
			errors.ErrorHandler(w, "Update user error", http.StatusInternalServerError, err)
			return
		}
	case http.MethodGet:
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
			errors.ErrorHandler(w, "GetUserById error", http.StatusInternalServerError, err)
			return
		}
		photo, err := s.Photo.GetPhotoByUser(*user.PhotoID)
		if err != nil {
			errors.ErrorHandler(w, "GetPhotoByUser error", http.StatusInternalServerError, err)
			return
		}
		w.Header().Add("Content-Type", "multipart/form-data")
		if _, err := w.Write([]byte(photo)); err != nil {
			return
		}
	default:
		errors.ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed, nil)
	}
}

func (s *service) UserSecurity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getUserSecurity(w, r)
	case http.MethodPut:
		s.putUserSecurity(w, r)
	default:
		errors.ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed, nil)
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
	if err := s.User.CreateUser(user); err != nil {
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
	err := s.Sm.Delete(r.Context())
	if err != nil {
		errors.ErrorHandler(w, "Error delete session", http.StatusInternalServerError, err)
		return
	}
	w.Header().Del("session_id")
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
	sessionId, err := s.Sm.Create(user)
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

func (s *service) securityResponse(w http.ResponseWriter, status int, respMessage string, err error) {
	w.WriteHeader(status)
	b, err := json.Marshal(&ResponseSecurity{
		Message: respMessage,
		Error:   err,
	})
	if _, err := w.Write([]byte(b)); err != nil {
		return
	}
}

func (s *service) CheckSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieSessionID, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			errors.ErrorHandler(w, "No session_id", http.StatusUnauthorized, err)
			return
		} else if err != nil {
			errors.ErrorHandler(w, "Error cookie", http.StatusUnauthorized, err)
			return
		}
		ctx := context.WithValue(r.Context(), "session_id", cookieSessionID.Value)
		if err := s.Sm.Check(&ctx); err != nil {
			errors.ErrorHandler(w, "Error check session", http.StatusUnauthorized, err)
			return
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
