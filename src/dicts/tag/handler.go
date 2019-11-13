package tag

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/cardTags"
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
	tagService      HandlerTagService
	cardTagsService cardTags.HandlerCardTagsService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewTagHandler(e *echo.Echo, userService user.HandlerUserService,
	tagService HandlerTagService,
	cardTagsService cardTags.HandlerCardTagsService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		userService:     userService,
		tagService:      tagService,
		cardTagsService: cardTagsService,
		securityService: securityService,
	}
	e.GET("/tags/:id", handler.Get)
	e.GET("/tags", handler.Fetch)
	e.POST("/card/:card_id/tags", handler.Post)
	e.PUT("/tags/:id", handler.Put)
	e.DELETE("/tags/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	data, err := h.tagService.FindTagByID(customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetTagById error", http.StatusBadRequest, err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h Handler) Fetch(ctx echo.Context) error {
	params := utils.ParseRequestParams(*ctx.Request().URL)
	tags, err := h.tagService.FetchTags(params.Limit, params.Offset)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Fetch error ", http.StatusBadRequest, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, tags)
}

func (h Handler) Post(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	tag, err := h.tagService.CreateTag(body)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "CreateTag error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}

	_, err = h.cardTagsService.CreateCardTags(tag.ID, customType.StringUUID(ctx.Param("card_id")))
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "CreateCardTags error", http.StatusInternalServerError, err)
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, tag)
}

func (h Handler) Put(ctx echo.Context) error {
	body, err := h.ReadBody(ctx.Request().Body)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Invalid body error", http.StatusInternalServerError, err)
		return err
	}
	tag, err := h.tagService.UpdateTag(body, customType.StringUUID(ctx.Param("id")))
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update tag error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, tag)
}

func (h Handler) Delete(ctx echo.Context) error {
	if err := h.tagService.DeleteTag(customType.StringUUID(ctx.Param("id"))); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Delete tag error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
