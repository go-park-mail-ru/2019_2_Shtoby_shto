package security

import (
	"2019_2_Shtoby_shto/src/customType"
	"net/http"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreateInstance(t *testing.T) {
	type args struct {
		sm *SessionManager
	}
	tests := []struct {
		name string
		args args
		want HandlerSecurity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInstance(tt.args.sm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteSession(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteSession(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_CreateSession(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		userID customType.StringUUID
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CreateSession(tt.args.w, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_checkNotSecurity(t *testing.T) {
	type args struct {
		route string
	}
	tests := []struct {
		name string
		s    service
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.checkNotSecurity(tt.args.route); got != tt.want {
				t.Errorf("service.checkNotSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CheckSession(t *testing.T) {
	type args struct {
		h echo.HandlerFunc
	}
	tests := []struct {
		name string
		s    *service
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CheckSession(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CheckSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
