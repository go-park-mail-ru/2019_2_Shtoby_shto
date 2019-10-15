package user

import (
	transport "2019_2_Shtoby_shto/src/handle"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userService HandlerUserService
	transport.HandlerImpl
}

func NewUserHandler(e *echo.Echo, userService HandlerUserService) {
	handler := UserHandler{
		userService: userService,
	}
	e.GET("/users", handler.Fetch)
	e.GET("/users/:id", handler.Get)
	e.POST("/users", handler.Post)
	e.PUT("/users/:id", handler.Put)
	e.DELETE("/users/:id", handler.Delete)
}

func (u UserHandler) Get(c echo.Context) error {
	return nil
}
