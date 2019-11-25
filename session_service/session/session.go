package session

import (
	"2019_2_Shtoby_shto/src/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

// Обработчик сессий
type SessionHandler interface {
	Create(context.Context, *UserID) (*Session, error)
	Check(context.Context, *SessionInfo) (*Session, error)
	Delete(context.Context, *SessionId) (*Nothing, error)
}

////easyjson:json
//type Session struct {
//	ID        StringUUID `json:"session_id"`
//	UserID    StringUUID `json:"user_id"`
//	CsrfToken string     `json:"csrf_token"`
//}

type SessionManager struct {
	cache *redis.Client
}

func NewSessionManager(addr, password string, dbNumber int) SessionHandler {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // use default Addr
		Password: password, // no password set
		DB:       dbNumber, // use default DB 0
	})
	return &SessionManager{
		cache: rdb,
	}
}

func (sm SessionManager) Create(ctx context.Context, user *UserID) (*Session, error) {
	//if !userID.IsUUID() {
	//	return nil, errors.New("userID is not uuid")
	//}
	id, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	HMACHashToken, err := NewHMACHashToken("1111")
	if err != nil {
		return nil, err
	}
	token, err := HMACHashToken.Create(id.String(), user.UserID, time.Now().AddDate(0, 0, 7).Unix())
	if err != nil {
		return nil, err
	}
	session := &Session{
		ID:        id.String(),
		UserID:    user.UserID,
		CsrfToken: token,
	}
	sessData, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}
	return session, sm.putSession(session.ID, sessData)
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

func (sm *SessionManager) Delete(ctx context.Context, sessionID *SessionId) (*Nothing, error) {
	if sessionID.SessionID == "" {
		return &Nothing{}, errors.New("Error session ")
	}

	return &Nothing{}, sm.cache.Del(sessionID.SessionID).Err()
}

func (sm *SessionManager) Check(ctx context.Context, session *SessionInfo) (*Session, error) {
	s := &Session{}
	sessionInfo, err := sm.getSession(session.ID)
	if err != nil {
		return nil, err
	}
	if sessionInfo == "" {
		return nil, errors.New("Missing session info")
	}
	if err := json.Unmarshal([]byte(sessionInfo), s); err != nil {
		return nil, err
	}
	HMACHashToken, err := NewHMACHashToken("1111")
	if err != nil {
		return nil, err
	}
	_, err = HMACHashToken.Check(s.ID, session.UserID, s.CsrfToken)
	if err != nil {
		return nil, err
	}
	return s, nil
}
