package user

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"bytes"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService       HandlerUserService
	boardUsersService boardUsers.HandlerBoardUsersService
	cardUsersService  сardUsers.HandlerCardUsersService
	securityService   security.HandlerSecurity
	handle.HandlerImpl
}

func NewUserHandler(e *echo.Echo, userService HandlerUserService,
	boardUsersService boardUsers.HandlerBoardUsersService,
	cardUsersService сardUsers.HandlerCardUsersService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:       userService,
		boardUsersService: boardUsersService,
		cardUsersService:  cardUsersService,
		securityService:   securityService,
	}
	e.POST("/login", handler.Login)
	e.GET("/logout", handler.Logout)
	e.GET("/users/all", handler.Fetch)
	e.GET("/users/:id", handler.Get)
	e.POST("/users/registration", handler.Post)
	e.PUT("/users", handler.Put)
	e.DELETE("/users", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	userID := customType.StringUUID(ctx.Param("id"))
	user, err := h.userService.GetUserById(userID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	users, err := h.userService.FetchUsers(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, users)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	user, err := h.userService.CreateUser(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "User not valid", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	ctx.Set("user_id", user.ID)
	h.SecurityResponse(ctx.Response(), http.StatusOK, "Registration is success, user id: "+user.ID.String(), nil)
	return err
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
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	user, err := h.userService.GetUserByLogin(body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Please, reg yourself", http.StatusUnauthorized, err)
		return err
	}

	if err := h.securityService.CreateSession(&ctx, user.ID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Set("user_id", user.ID)
	return ctx.JSON(http.StatusOK, user)
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

func (h Handler) Delete(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("get user_id failed"))
		return errors.New("get user_id failed")
	}
	if err := h.userService.DeleteUser(userID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "DeleteUser failed", http.StatusInternalServerError, errors.New("DeleteUser failed"))
		return err
	}
	return ctx.JSON(http.StatusOK, userID)
}
