package security

import (
	"2019_2_Shtoby_shto/session_service/session"
	. "2019_2_Shtoby_shto/src/customType"
	"context"
	"errors"
)

// Обработчик сессий
type SessionHandler interface {
	Create(userID StringUUID) (*session.Session, error)
	Check(sessionID, userID string) (*session.Session, error)
	Delete(sessionID string) error
}

type SessionManager struct {
	sessService *session.SecurityClient
}

func NewSessionManager(sessService *session.SecurityClient) SessionHandler {
	return &SessionManager{
		sessService: sessService,
	}
}

func (sm SessionManager) Create(userID StringUUID) (*session.Session, error) {
	if !userID.IsUUID() {
		return nil, errors.New("userID is not uuid")
	}
	// Create session
	create := &session.UserID{
		UserID: userID.String(),
	}

	sess, err := (*sm.sessService).Create(context.Background(), create)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (sm *SessionManager) Delete(sessionID string) error {
	if sessionID == "" {
		return errors.New("Error session ")
	}
	del := &session.SessionId{
		SessionID: sessionID,
	}
	_, err := (*sm.sessService).Delete(context.Background(), del)
	if err != nil {
		return err
	}
	return nil
}

func (sm *SessionManager) Check(sessionID, userID string) (*session.Session, error) {
	sessInfo := &session.SessionInfo{
		ID:     sessionID,
		UserID: userID,
	}
	s, err := (*sm.sessService).Check(context.Background(), sessInfo)
	if err != nil {
		return nil, err
	}
	return s, nil
}
