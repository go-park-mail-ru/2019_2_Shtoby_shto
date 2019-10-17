package user

import (
	"2019_2_Shtoby_shto/src/customType"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	handle "2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"encoding/json"
	"errors"
	"github.com/labstack/echo"
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
	e.GET("/logout", handler.securityService.CheckSession(handler.Logout))
	e.GET("/users", handler.securityService.CheckSession(handler.Fetch))
	e.GET("/users/:id", handler.securityService.CheckSession(handler.Get))
	e.POST("/users", handler.securityService.CheckSession(handler.Post))
	e.PUT("/users/:id", handler.securityService.CheckSession(handler.Put))
	e.DELETE("/users/:id", handler.securityService.CheckSession(handler.Delete))
}

func (h Handler) Get(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		errorsLib.ErrorHandler(ctx.Response(), "get user_id failed", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("get user_id failed")
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
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
	user := User{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}
	id, err := utils.GenerateUUID()
	if err != nil || id.String() == "" {
		errorsLib.ErrorHandler(ctx.Response(), "Error create UUID", http.StatusInternalServerError, err)
		return err
	}
	user.ID = customType.StringUUID(id.String())
	if err := h.userService.CreateUser(user); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "User not valid", http.StatusBadRequest, err)
		return err
	}
	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		return err
	}
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Registration is success", err)
	return nil
}

func (h Handler) Put(ctx echo.Context) error {
	user := User{}
	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("download fail")
	}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}
	if err := h.userService.UpdateUser(user, customType.StringUUID(userID)); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusBadRequest, err)
		return err
	}
	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Update is success", nil)
	return nil
}
func (h Handler) Login(ctx echo.Context) error {
	curUser := User{}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&curUser); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}

	user, err := h.userService.GetUserByLogin(curUser.Login)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Please, reg yourself", http.StatusUnauthorized, err)
		return err
	}

	if strings.Compare(user.Password, curUser.Password) != 0 {
		errorsLib.ErrorHandler(ctx.Response(), "Ne tot password )0))", http.StatusBadRequest, err)
		return err
	}

	if err := h.securityService.CreateSession(ctx.Response(), user.ID); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create session error", http.StatusInternalServerError, err)
		return err
	}

	ctx.Set("user_id", user.ID)

	h.securityService.SecurityResponse(ctx.Response(), http.StatusOK, "Login", err)
	return nil
}

func (h Handler) Logout(ctx echo.Context) error {
	return nil
}
