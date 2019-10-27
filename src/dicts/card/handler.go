package card

import (
	"2019_2_Shtoby_shto/src/customType"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService      user.HandlerUserService
	cardService      HandlerCardService
	cardUsersService сardUsers.HandlerCardUsersService
	securityService  security.HandlerSecurity
	handle.HandlerImpl
}

func NewCardHandler(e *echo.Echo, userService user.HandlerUserService, cardService HandlerCardService, cardUsersService сardUsers.HandlerCardUsersService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:      userService,
		cardService:      cardService,
		cardUsersService: cardUsersService,
		securityService:  securityService,
	}
	e.GET("/cards/:id", handler.Get)
	e.POST("/cards", handler.Post)
	e.PUT("/cards/:id", handler.Put)
	e.DELETE("/cards/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.cardService.FindCardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetCardById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	responseData, err := h.cardService.CreateCard(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, responseData)
}

func (h Handler) Put(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	board, err := h.cardService.UpdateCard(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update card error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.cardService.DeleteCard(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete card error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
