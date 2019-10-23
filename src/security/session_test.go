package security

import (
	. "2019_2_Shtoby_shto/src/customType"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewSessionManager(t *testing.T) {
	type args struct {
		addr     string
		password string
		dbNumber int
	}
	tests := []struct {
		name string
		args args
		want *SessionManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSessionManager(tt.args.addr, tt.args.password, tt.args.dbNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSessionManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_Create(t *testing.T) {
	type args struct {
		userID StringUUID
	}
	tests := []struct {
		name    string
		sm      SessionManager
		args    args
		want    *SessionID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sm.Create(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionManager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionManager_putSession(t *testing.T) {
	type args struct {
		sessionID StringUUID
		userID    StringUUID
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
			if err := tt.sm.putSession(tt.args.sessionID, tt.args.userID); (err != nil) != tt.wantErr {
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
			got, err := tt.sm.getSession(tt.args.cacheID)
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
			if err := tt.sm.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionManager_Check(t *testing.T) {
	type args struct {
		ctx *echo.Context
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
			if err := tt.sm.Check(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SessionManager.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
