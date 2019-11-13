package comment

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
	userService     user.HandlerUserService
	commentService  HandlerCommentService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewCommentHandler(e *echo.Echo, userService user.HandlerUserService,
	commentService HandlerCommentService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		commentService:  commentService,
		securityService: securityService,
	}
	e.GET("/comments/:id", handler.Get)
	e.GET("/comments", handler.Fetch)
	e.POST("/comments", handler.Post)
	e.PUT("/comments/:id", handler.Put)
	e.DELETE("/comments/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.commentService.FindCommentByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetCommentById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	comments, err := h.commentService.FetchComments(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, comments)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	responseData, err := h.commentService.CreateComment(body)
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
	comment, err := h.commentService.UpdateComment(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update comment error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, comment)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.commentService.DeleteComment(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete comment error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
