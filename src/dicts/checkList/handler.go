package checkList

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService      user.HandlerUserService
	checkListService HandlerCheckListService
	securityService  security.HandlerSecurity
	handle.HandlerImpl
}

func NewCheckListHandler(e *echo.Echo, userService user.HandlerUserService,
	checkListService HandlerCheckListService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:      userService,
		checkListService: checkListService,
		securityService:  securityService,
	}
	e.GET("/checkLists/:id", handler.Get)
	e.GET("/checkLists", handler.Fetch)
	e.POST("/checkLists", handler.Post)
	e.PUT("/checkLists/:id", handler.Put)
	e.DELETE("/checkLists/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.checkListService.FindCheckListByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetCheckListById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	checkLists, err := h.checkListService.FetchCheckLists(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, checkLists)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	responseData, err := h.checkListService.CreateCheckList(body)
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
	checkList, err := h.checkListService.UpdateCheckList(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update checkList error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, checkList)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.checkListService.DeleteCheckList(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete checkList error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
