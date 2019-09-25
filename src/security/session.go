package security

import (
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/utils"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

// Обработчик сессий
type SessionHandler interface {
	Create(user user.User) (*SessionID, error)
	Check(in *SessionID) (bool, error)
	Delete(in *SessionID) error
}

// Описание сессии
//type Session struct {
//	Login    string
//	Password string
//}

// id сессии
type SessionID struct {
	ID string `json:"session_id"`
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

func (sm SessionManager) Create(user user.User) (*SessionID, error) {
	expire := time.Duration(24 * time.Hour)
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		return nil, err
	}
	session := SessionID{id.String()}
	return &session, sm.putSession(session.ID, user, expire)
}

func (sm *SessionManager) putSession(id string, user user.User, expire time.Duration) error {
	return sm.cache.Set(id, user.Id, 0).Err()
}

func (sm *SessionManager) getSession(idIn string) error {
	val, err := sm.cache.Get(idIn).Result()
	if err == redis.Nil {
		return errors.New("missing_key does not exist")
	}
	if err != nil {
		return err
	}
	println(val)
	return nil
}

func (sm *SessionManager) Check(sessionId *SessionID) (bool, error) {
	return true, sm.getSession(sessionId.ID)
}

func (sm *SessionManager) Delete(in *SessionID) error {
	return sm.cache.Del(in.ID).Err()
}
