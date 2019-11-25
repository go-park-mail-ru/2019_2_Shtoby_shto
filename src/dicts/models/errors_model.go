package models

// Описание структуры ответа при ошибке
//easyjson:json
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
