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
	board, err := h.boardService.FindBoardByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusBadRequest, err)
		return err
	}
	ctx.Response().WriteHeader(http.StatusOK)
	ctx.Response().Header().Set("Content-Type", "application/json")
	b, err := board.MarshalJSON()
	if _, err := ctx.Response().Write([]byte(b)); err != nil {
		ctx.Logger().Error(err)
		return err
	}
	return nil
}

func (h Handler) Post(ctx echo.Context) error {
	board := &Board{}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		return err
	}
	if err := board.UnmarshalJSON(buf.Bytes()); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	if err := h.boardService.CreateBoard(board); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Create error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	return nil
}

func (h Handler) Put(ctx echo.Context) error {
	board := &Board{}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		return err
	}
	if err := board.UnmarshalJSON(buf.Bytes()); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Decode error", http.StatusInternalServerError, err)
		return err
	}
	if err := h.boardService.UpdateBoard(board, customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.boardService.DeleteBoard(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "UpdateBoard error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
