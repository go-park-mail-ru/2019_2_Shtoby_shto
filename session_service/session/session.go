package session

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/utils"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

// Обработчик сессий
type SessionHandler interface {
	Create(userID StringUUID) (*Session, error)
	Check(sessionID string) (*Session, error)
	Delete(sessionID string) error
}

//easyjson:json
type Session struct {
	ID        StringUUID `json:"session_id"`
	UserID    StringUUID `json:"user_id"`
	CsrfToken string     `json:"csrf_token"`
}

type SessionManager struct {
	cache *redis.Client
}

func NewSessionManager(addr, password string, dbNumber int) *SessionManager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // use default Addr
		Password: password, // no password set
		DB:       dbNumber, // use default DB 0
	})
	return &SessionManager{
		cache: rdb,
	}
}

func (sm SessionManager) Create(userID StringUUID) (*Session, error) {
	if !userID.IsUUID() {
		return nil, errors.New("userID is not uuid")
	}
	id, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	HMACHashToken, err := NewHMACHashToken("1111")
	if err != nil {
		return nil, err
	}
	token, err := HMACHashToken.Create(id.String(), userID.String(), time.Now().AddDate(0, 0, 7).Unix())
	if err != nil {
		return nil, err
	}
	session := &Session{
		ID:        StringUUID(id.String()),
		UserID:    userID,
		CsrfToken: token,
	}
	sessData, err := session.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return session, sm.putSession(session.ID.String(), sessData)
}

func (sm *SessionManager) putSession(sessionID string, sessionInfo []byte) error {
	//todo::set expire
	//expire := time.Duration(24 * time.Hour)
	return sm.cache.Set(sessionID, sessionInfo, 0).Err()
}

func (sm *SessionManager) getSession(cacheID string) (string, error) {
	val, err := sm.cache.Get(cacheID).Result()
	if err == redis.Nil {
		return "", errors.New("missing_key does not exist")
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (sm *SessionManager) Delete(sessionID string) error {
	if sessionID == "" {
		return errors.New("Error session ")
	}

	return sm.cache.Del(sessionID).Err()
}

func (sm *SessionManager) Check(sessionID string) (*Session, error) {
	s := &Session{}
	sessionInfo, err := sm.getSession(sessionID)
	if err != nil {
		return nil, err
	}
	if sessionInfo == "" {
		return nil, errors.New("Missing session info")
	}
	if err := s.UnmarshalJSON([]byte(sessionInfo)); err != nil {
		return nil, err
	}
	HMACHashToken, err := NewHMACHashToken("1111")
	if err != nil {
		return nil, err
	}
	_, err = HMACHashToken.Check(sessionID, s.UserID.String(), s.CsrfToken)
	if err != nil {
		return nil, err
	}
	return s, nil
}
