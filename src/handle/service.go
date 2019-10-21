// TODO:: work in progress
package handle

import (
	"2019_2_Shtoby_shto/src/dicts"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Описание структуры ответа на list-запрос
type FetchResult struct {
	Name  string   `json:"name"`
	Total int      `json:"total"`
	Items []string `json:"items"`
}

type ResponseSecurity struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Handler interface {
	// Get
	Get(c echo.Context)
	// Fetch all
	Fetch(c echo.Context)
	// Patch
	//Patch(w http.ResponseWriter, req *http.Request)
	// Post
	Post(c echo.Context)
	// Delete ...handlerManager
	Delete(c echo.Context)
	// Put
	Put(c echo.Context)

	CreateInstance() dicts.Dict
}

// Класс реализующий транспортный уровень
type HandlerImpl struct {
	Handler
}

func (h *HandlerImpl) SecurityResponse(w http.ResponseWriter, status int, respMessage string, err error) {
	w.WriteHeader(status)
	b, err := json.Marshal(&ResponseSecurity{
		Message: respMessage,
		Error:   err,
	})
	if _, err := w.Write([]byte(b)); err != nil {
		return
	}
}

func (h HandlerImpl) CreateInstance() dicts.Dict {
	return &dicts.BaseInfo{}
}

// Get
func (h HandlerImpl) Get(c echo.Context) error {
	return nil
}

func (h HandlerImpl) Fetch(c echo.Context) error {
	return nil
}

// Post
func (h HandlerImpl) Post(c echo.Context) error {
	return nil
}

// Patch
func (h HandlerImpl) Put(c echo.Context) error {
	return nil
}

// Delete
func (h HandlerImpl) Delete(c echo.Context) error {
	return nil
}
