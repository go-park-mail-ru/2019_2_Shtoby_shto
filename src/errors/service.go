package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/log"
)

// Описание структуры ответа при ошибке
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ErrorHandler(response *echo.Response, message string, status int, err error) {
	errorMessage := "Error!"
	if err != nil {
		errorMessage = err.Error()
	}
	resp := &ErrorResponse{
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
