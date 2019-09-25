package security

import (
	"2019_2_Shtoby_shto/src/utils"
	"github.com/go-redis/redis"
)

// Обработчик сессий
type SessionHandler interface {
	Create(in *Session) (*SessionID, error)
	Check(in *SessionID) (bool, error)
	Delete(in *SessionID)
}

// Описание сессии
type Session struct {
	Login    string
	Password string
}

// id сессии
type SessionID struct {
	ID string `json:"session_id"`
}

// TODO::Redis
type SessionManager struct {
	cache *redis.Client
}

func NewSessionManager() SessionHandler {
	return &SessionManager{
		cache: &redis.Client{},
	}
}

func (sm *SessionManager) Create(in *Session) (*SessionID, error) {
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		return nil, err
	}
	session := SessionID{id.String()}
	// putSession(session.ID, user.Id)
	return &session, err
}

func (sm *SessionManager) putSession(id, userId string) error {

	return nil
}

func (sm *SessionManager) getSession(idIn string) error {
	return nil
}

func (sm *SessionManager) Check(sessionId *SessionID) (bool, error) {
	return true, nil
}

// TODO:: удаление сессий
func (sm *SessionManager) Delete(in *SessionID) {

}
