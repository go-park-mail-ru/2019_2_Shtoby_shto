package photo

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"bufio"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	photoService    HandlerPhotoService
	userService     user.HandlerUserService
	securityService security.HandlerSecurity
	handle.HandlerImpl
}

func NewPhotoHandler(e *echo.Echo, photoService HandlerPhotoService,
	userService user.HandlerUserService,
	securityService security.HandlerSecurity) {
	handler := Handler{
		photoService:    photoService,
		userService:     userService,
		securityService: securityService,
	}
	e.GET("/photo/:id", handler.Get)
	e.POST("/photo", handler.Post)
	e.PUT("/photo/:id", handler.Put)
	e.DELETE("/photo/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	photoID := customType.StringUUID(ctx.Param("id"))
	photo, err := h.photoService.GetPhotoByUser(photoID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetPhotoByUser error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Add("Content-Type", "multipart/form-data")
	if _, err := ctx.Response().Write([]byte(photo)); err != nil {
		ctx.Logger().Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, photo)
}

func (h Handler) Post(ctx echo.Context) error {
	rr := bufio.NewReader(ctx.Request().Body)
	photo, err := h.photoService.DownloadPhoto(rr)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, err)
		return err
	}

	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		ctx.Logger().Error("error get user_id")
		errorsLib.ErrorHandler(ctx.Response(), "error get user_id", http.StatusInternalServerError, errors.New("error get user_id"))
		return errors.New("error get user_id")
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusInternalServerError, err)
		return err
	}
	user.PhotoID = &photo.ID
	updateUser, err := user.MarshalJSON()
	if err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusInternalServerError, err)
		return err
	}
	if err := h.userService.UpdateUser(updateUser, userID); err != nil {
		ctx.Logger().Error(err)
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusInternalServerError, err)
		return err
	}
	return ctx.JSON(http.StatusOK, photo)
}
