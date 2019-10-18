package errors

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
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
	b, _ := json.Marshal(&ErrorResponse{
		Status:  status,
		Message: message,
		Error:   errorMessage,
	})
	response.WriteHeader(status)
	if _, err := response.Write([]byte(b)); err != nil {

	}
}
