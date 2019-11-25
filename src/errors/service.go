package errors

import (
	"2019_2_Shtoby_shto/src/dicts/models"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/log"
)

func ErrorHandler(response *echo.Response, message string, status int, err error) {
	errorMessage := "Error!"
	if err != nil {
		errorMessage = err.Error()
	}
	resp := &models.ErrorResponse{
		Status:  status,
		Message: message,
		Error:   errorMessage,
	}
	r, err := resp.MarshalJSON()
	response.WriteHeader(status)
	if err != nil {
		log.Error(err)
		return
	}
	if _, err := response.Write(r); err != nil {
		log.Error(err)
	}
}
