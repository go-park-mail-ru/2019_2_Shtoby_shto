package user

import (
	"2019_2_Shtoby_shto/src/customType"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	handle "2019_2_Shtoby_shto/src/handle"
	"encoding/json"
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	userService HandlerUserService
	handle.HandlerImpl
}

func NewUserHandler(e *echo.Echo, userService HandlerUserService) {
	handler := Handler{
		userService: userService,
	}
	e.GET("/users", handler.Fetch)
	e.GET("/users/:id", handler.Get)
	e.POST("/users", handler.Post)
	e.PUT("/users/:id", handler.Put)
	e.DELETE("/users/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("download fail")
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusBadRequest, err)
		return err
	}
	ctx.Response().WriteHeader(http.StatusOK)
	b, err := json.Marshal(&user)
	if _, err := ctx.Response().Write([]byte(b)); err != nil {
		return err
	}
	return nil
}

func (h Handler) Post(ctx echo.Context) error {

	return nil
}
