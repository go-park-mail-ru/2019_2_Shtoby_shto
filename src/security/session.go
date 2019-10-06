package security

import (
	. "2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/utils"
	"context"
	"errors"
	"github.com/go-redis/redis"
)

// Обработчик сессий
type SessionHandler interface {
	Create(user user.User) (*SessionID, error)
	Check(ctx *context.Context) error
	Delete(ctx context.Context) error
}

// id сессии
type SessionID struct {
	ID StringUUID `json:"session_id"`
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
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		return nil, err
	}
	session := SessionID{StringUUID(id.String())}
	return &session, sm.putSession(session.ID, user)
}

func (sm *SessionManager) putSession(id StringUUID, user user.User) error {
	//todo::set expire
	//expire := time.Duration(24 * time.Hour)
	return sm.cache.Set(id.String(), user.ID.String(), 0).Err()
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

func (sm *SessionManager) Delete(ctx context.Context) error {
	return sm.cache.Del(ctx.Value("session_id").(string)).Err()
}

func (sm *SessionManager) Check(ctx *context.Context) error {
	userId, err := sm.getSession((*ctx).Value("session_id").(string))
	if err != nil {
		return err
	}
	if userId == "" {
		return errors.New("Missing userId")
	}
	*ctx = context.WithValue(*ctx, "user_id", StringUUID(userId))
	return nil
}
