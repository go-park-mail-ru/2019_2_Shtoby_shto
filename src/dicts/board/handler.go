package board

import (
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	userService     user.HandlerUserService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewBoardHandler(e *echo.Echo, userService user.HandlerUserService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		securityService: securityService,
	}
	e.GET("/board", handler.Get)
	e.POST("/board", handler.Post)
	e.PUT("/board/:id", handler.Put)
	e.DELETE("/board/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	return nil
}

func (h Handler) Post(ctx echo.Context) error {
	return nil
}

func (h Handler) Put(ctx echo.Context) error {
	return nil
}

func (h Handler) Delete(ctx echo.Context) error {
	return nil
}
