package photo

import (
	transport "2019_2_Shtoby_shto/src/handle"
	"github.com/labstack/echo"
)

type PhotoHandler struct {
	PhotoService HandlerPhotoService
	transport.HandlerImpl
}

func NewUserHandler(e *echo.Echo, photoService HandlerPhotoService) {
	handler := PhotoHandler{
		PhotoService: photoService,
	}
	e.GET("/photo", handler.Get)
	e.PUT("/photo/:id", handler.Put)
	e.DELETE("/photo/:id", handler.Delete)
}

func (u PhotoHandler) Get(c echo.Context) error {
	return nil
}
