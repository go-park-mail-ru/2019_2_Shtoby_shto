package transport

import (
	"2019_2_Shtoby_shto/src/dicts"
	"2019_2_Shtoby_shto/src/errors"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Описание структуры ответа на list-запрос
type FetchResult struct {
	Name  string   `json:"name"`
	Total int      `json:"total"`
	Items []string `json:"items"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Handler interface {
	// Get
	Get(w http.ResponseWriter, req *http.Request)
	// Patch
	//Patch(w http.ResponseWriter, req *http.Request)
	// Post
	Post(w http.ResponseWriter, req *http.Request)
	// Delete ...handlerManager
	Delete(w http.ResponseWriter, req *http.Request)
	// Put
	Put(w http.ResponseWriter, req *http.Request)

	CreateInstance() dicts.Dict
}

// Класс реализующий транспортный уровень

type HandlerImpl struct {
	Handler
}

func (h HandlerImpl) CreateInstance() dicts.Dict {
	return &dicts.BaseInfo{}
}

// Get
func (h HandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id != "" {
		if err := h.fetchOne(id, w); err != nil {
			return
		}
	} else {
		if err := h.fetchList(w); err != nil {
			return
		}
	}
}

func (h HandlerImpl) fetchList(w http.ResponseWriter) error {
	var err error

	if err != nil {
		errors.ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError, err)
		return err
	}
	b, _ := json.Marshal(&FetchResult{
		Name: "List",
		//Total: count,
		Items: nil,
	})
	w.Write([]byte(b))
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h HandlerImpl) fetchOne(id string, w http.ResponseWriter) error {
	return nil
}

// Post
func (h HandlerImpl) Post(w http.ResponseWriter, r *http.Request) {

}

// Patch
func (h HandlerImpl) Put(w http.ResponseWriter, r *http.Request) {

}

// Delete
func (h HandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {

}
