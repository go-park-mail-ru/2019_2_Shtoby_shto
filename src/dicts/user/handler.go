package user

import (
	"2019_2_Shtoby_shto/src/customType"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"bytes"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService     HandlerUserService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewUserHandler(e *echo.Echo, userService HandlerUserService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		securityService: securityService,
	}
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout)
	e.GET("/users", handler.Fetch)
	e.GET("/users/:id", handler.Get)
	e.POST("/users/registration", handler.Post)
	e.PUT("/users/:id", handler.Put)
	e.DELETE("/users/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("get user_id failed")
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

func (h Handler) Post(ctx echo.Context) error {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body", http.StatusInternalServerError, err)
		return err
	}
	user, err := h.userService.CreateUser(buf.Bytes())
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "User not valid", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	h.SecurityResponse(ctx.Response(), http.StatusOK, "Registration is success, user id: "+user.ID.String(), nil)
	return nil
}

func (h Handler) Put(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("get user_id failed"))
		return errors.New("get user_id failed")
	}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		return err
	}
	if err := h.userService.UpdateUser(buf.Bytes(), customType.StringUUID(userID)); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusBadRequest, err)
		return err
	}
	h.SecurityResponse(ctx.Response(), http.StatusOK, "Update is success", nil)
	return nil
}

func (h Handler) Login(ctx echo.Context) error {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body", http.StatusInternalServerError, err)
		return err
	}
	user, err := h.userService.GetUserByLogin(buf.Bytes())
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Please, reg yourself", http.StatusUnauthorized, err)
		return err
	}

	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Set("user_id", user.ID)
	h.SecurityResponse(ctx.Response(), http.StatusOK, "Login", err)
	return nil
}

func (h Handler) Logout(ctx echo.Context) (err error) {
	if err = h.securityService.DeleteSession(ctx); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Error delete session", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Del("session_id")
	h.SecurityResponse(ctx.Response(), http.StatusOK, "Logout", err)
	return err
}
