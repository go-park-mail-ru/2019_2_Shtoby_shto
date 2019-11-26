package session

import (
	. "2019_2_Shtoby_shto/src/customType"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSessionManager_Create(t *testing.T) {
	sm := NewSessionManager("localhost:6379", "", 0)
	tests := []struct {
		name    string
		userID  StringUUID
		wantErr bool
	}{
		{
			name:    "test 1",
			userID:  StringUUID("123"),
			wantErr: true,
		},
		{
			name:    "test 2",
			userID:  StringUUID("33b42c6b-6819-4254-b2e4-ee4b21fbbd10"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(UserID, tt.userID) {
				t.Errorf("SessionManager.Create() = %v, want %v", got, tt.userID)
			}
		})
	}
}

func TestSessionManager_putSession(t *testing.T) {
	sm := NewSessionManager("localhost:6379", "", 0)
	tests := []struct {
		name    string
		session Session
		wantErr bool
	}{
		{
			name: "test 1",
			session: Session{
				ID:        "11112c6b-6819-4254-b2e4-ee4b21fbbd10",
				UserID:    "33b42c6b-6819-4254-b2e4-ee4b21fbbd10",
				CsrfToken: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := MarshalJSON()
			if err != nil {
				t.Errorf("SessionManager.putSession() error unmarshal ")
			}
			if err := putSession(tt.session.ID.String(), data); (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.putSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionManager_getSession(t *testing.T) {
	type args struct {
		cacheID string
	}
	tests := []struct {
		name    string
		sm      *SessionManager
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSession(tt.args.cacheID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.getSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SessionManager.getSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_Delete(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		sm      *SessionManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionManager_Check(t *testing.T) {
	sm := NewSessionManager("localhost:6379", "", 0)

	type args struct {
		ctx *echo.Context
	}
	tests := []struct {
		name      string
		sessionID StringUUID
		wantErr   bool
	}{
		{
			name:      "test 1",
			sessionID: "b1f395db-ddf9-4629-824f-0bb81d53a57b",
			wantErr:   false,
		},
		{
			name:      "test 1",
			sessionID: "b1f395db-824f-0bb81d53a57b",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := Create(tt.sessionID)
			if err != nil {
				t.Errorf("Create error ")
			}
			checkS, err := Check(tt.sessionID.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
			if checkS != nil && !reflect.DeepEqual(ID, ID) {
				t.Errorf("Check() = %v, want %v", ID, ID)
			}
		})
	}
}
