package photo

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	errorsLib "2019_2_Shtoby_shto/src/errors"
	mainHandle "2019_2_Shtoby_shto/src/handle"
	"bufio"
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	PhotoService HandlerPhotoService
	UserService  user.HandlerUserService
	mainHandle.HandlerImpl
}

func NewPhotoHandler(e *echo.Echo, photoService HandlerPhotoService, userService user.HandlerUserService) {
	handler := Handler{
		PhotoService: photoService,
		UserService:  userService,
	}
	e.GET("/photo", handler.Get)
	e.POST("/photo", handler.Post)
	e.PUT("/photo/:id", handler.Put)
	e.DELETE("/photo/:id", handler.Delete)
}

func (h Handler) Get(ctx echo.Context) error {
	userID := ctx.Get("user_id").(customType.StringUUID)
	user, err := h.UserService.GetUserById(userID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusInternalServerError, err)
		return err
	}
	photo, err := h.PhotoService.GetPhotoByUser(*user.PhotoID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "GetPhotoByUser error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Add("Content-Type", "multipart/form-data")
	if _, err := ctx.Response().Write([]byte(photo)); err != nil {
		return err
	}
	return nil
}

func (h Handler) Post(ctx echo.Context) error {
	rr := bufio.NewReader(ctx.Request().Body)
	photoID, err := h.PhotoService.DownloadPhoto(rr)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, err)
		return err
	}

	userID, ok := ctx.Get("user_id").(customType.StringUUID)
	if !ok {
		errorsLib.ErrorHandler(ctx.Response(), "download fail", http.StatusInternalServerError, errors.New("download fail"))
		return errors.New("download fail")
	}

	user, err := h.UserService.GetUserById(userID)
	if err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusInternalServerError, err)
		return err
	}
	user.PhotoID = &photoID
	if err := h.UserService.UpdateUser(user, userID); err != nil {
		errorsLib.ErrorHandler(ctx.Response(), "Update user error", http.StatusInternalServerError, err)
		return err
	}
	return nil
}
