package user

import (
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		e               *echo.Echo
		userService     HandlerUserService
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
			NewUserHandler(tt.args.e, tt.args.userService, tt.args.securityService)
		})
	}
}

func TestHandler_Get(t *testing.T) {
	type fields struct {
		userService     HandlerUserService
		securityService security.HandlerSecurity
		HandlerImpl     handle.HandlerImpl
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				userService:     tt.fields.userService,
				securityService: tt.fields.securityService,
				HandlerImpl:     tt.fields.HandlerImpl,
			}
			if err := h.Get(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_Post(t *testing.T) {
	type fields struct {
		userService     HandlerUserService
		securityService security.HandlerSecurity
		HandlerImpl     handle.HandlerImpl
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				userService:     tt.fields.userService,
				securityService: tt.fields.securityService,
				HandlerImpl:     tt.fields.HandlerImpl,
			}
			if err := h.Post(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_Put(t *testing.T) {
	type fields struct {
		userService     HandlerUserService
		securityService security.HandlerSecurity
		HandlerImpl     handle.HandlerImpl
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				userService:     tt.fields.userService,
				securityService: tt.fields.securityService,
				HandlerImpl:     tt.fields.HandlerImpl,
			}
			if err := h.Put(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type fields struct {
		userService     HandlerUserService
		securityService security.HandlerSecurity
		HandlerImpl     handle.HandlerImpl
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				userService:     tt.fields.userService,
				securityService: tt.fields.securityService,
				HandlerImpl:     tt.fields.HandlerImpl,
			}
			if err := h.Login(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_Logout(t *testing.T) {
	type fields struct {
		userService     HandlerUserService
		securityService security.HandlerSecurity
		HandlerImpl     handle.HandlerImpl
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				userService:     tt.fields.userService,
				securityService: tt.fields.securityService,
				HandlerImpl:     tt.fields.HandlerImpl,
			}
			if err := h.Logout(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
