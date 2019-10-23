package board

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"bytes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	userService     user.HandlerUserService
	boardService    HandlerBoardService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewBoardHandler(e *echo.Echo, userService user.HandlerUserService, boardService HandlerBoardService, securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		boardService:    boardService,
		securityService: securityService,
	}
	e.GET("/board/:id", handler.Get)
	e.POST("/board", handler.Post)
	e.PUT("/board/:id", handler.Put)
	e.DELETE("/board/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.boardService.FindBoardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Post(ctx echo.Context) error {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	responseData, err := h.boardService.CreateBoard(buf.Bytes())
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, responseData)
}

func (h Handler) Put(ctx echo.Context) error {
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	board, err := h.boardService.UpdateBoard(buf.Bytes(), customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, board)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.boardService.DeleteBoard(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
