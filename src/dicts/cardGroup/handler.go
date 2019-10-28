package cardGroup

import (
	"2019_2_Shtoby_shto/src/customType"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	cardGroupService HandlerCardGroupService
	securityService  security.HandlerSecurity
	handle.HandlerImpl
}

func NewCardGroupHandler(e *echo.Echo, cardGroupService HandlerCardGroupService, securityService security.HandlerSecurity) {
	handler := Handler{
		cardGroupService: cardGroupService,
		securityService:  securityService,
	}
	e.GET("/card-group/:id", handler.Get)
	e.GET("/card-group", handler.Fetch)
	e.POST("/card-group", handler.Post)
	e.PUT("/card-group/:id", handler.Put)
	e.DELETE("/card-group/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.cardGroupService.FindCardGroupByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetCardGroupById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	users, err := h.cardGroupService.FetchCardGroup(params.Limit, params.Offset)
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
	responseData, err := h.cardGroupService.CreateCardGroup(body)
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
	board, err := h.cardGroupService.UpdateCardGroup(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update card error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.cardGroupService.DeleteCardGroup(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete card error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
