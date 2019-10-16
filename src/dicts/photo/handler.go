package photo

import (
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/errors"
	mainHandle "2019_2_Shtoby_shto/src/handle"
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
	var userID = ctx.Get("user_id").(customType.StringUUID)
	user, err := h.UserService.GetUserById(userID)
	if err != nil {
		errors.ErrorHandler(ctx.Response(), "GetUserById error", http.StatusInternalServerError, err)
		return err
	}
	photo, err := h.PhotoService.GetPhotoByUser(*user.PhotoID)
	if err != nil {
		errors.ErrorHandler(ctx.Response(), "GetPhotoByUser error", http.StatusInternalServerError, err)
		return err
	}
	ctx.Response().Header().Add("Content-Type", "multipart/form-data")
	if _, err := ctx.Response().Write([]byte(photo)); err != nil {
		return err
	}
	return nil
}
