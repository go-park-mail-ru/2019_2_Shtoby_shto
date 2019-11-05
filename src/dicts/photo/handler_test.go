package photo

import (
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/security"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewPhotoHandler(t *testing.T) {
	type args struct {
		e               *echo.Echo
		photoService    HandlerPhotoService
		userService     user.HandlerUserService
		securityService security.HandlerSecurity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewPhotoHandler(tt.args.e, tt.args.photoService, tt.args.userService, tt.args.securityService)
		})
	}
}

func TestHandler_Get(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		h       Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Get(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_Post(t *testing.T) {
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		h       Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Post(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
