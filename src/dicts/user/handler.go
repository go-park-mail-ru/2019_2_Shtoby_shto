package user

import (
	"2019_2_Shtoby_shto/src/customType"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
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
	ctx.Response().WriteHeader(http.StatusOK)
	ctx.Response().Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(&user)
	if _, err := ctx.Response().Write([]byte(b)); err != nil {
		ctx.Logger().Error(err)
		return err
	}
	return nil
}

func (h Handler) Post(ctx echo.Context) error {
	user := User{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		errorsLib.ErrorHandler(ctx.Response(), "Error create UUID", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	user.ID = customType.StringUUID(id.String())
	if err := h.userService.CreateUser(user); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "User not valid", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Registration is success", err)
	return nil
}

func (h Handler) Put(ctx echo.Context) error {
	user := User{}
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("get user_id failed")
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("get user_id failed"))
		return errors.New("get user_id failed")
	}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}
	if err := h.userService.UpdateUser(user, customType.StringUUID(userID)); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusBadRequest, err)
		return err
	}
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Update is success", nil)
	return nil
}
func (h Handler) Login(ctx echo.Context) error {
	curUser := User{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&curUser); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}

	user, err := h.userService.GetUserByLogin(curUser.Login)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Please, reg yourself", http.StatusUnauthorized, err)
		return err
	}
	if strings.Compare(user.Password, curUser.Password) != 0 {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Ne tot password )0))", http.StatusBadRequest, err)
		return err
	}

	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Set("user_id", user.ID)
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Login", err)
	return nil
}

func (h Handler) Logout(ctx echo.Context) (err error) {
	if err = h.securityService.Logout(ctx); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Error delete session", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Del("session_id")
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Logout", err)
	return err
}
