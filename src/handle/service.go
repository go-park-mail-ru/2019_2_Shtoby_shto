package transport

import (
	"2019_2_Shtoby_shto/src/errors"
	"2019_2_Shtoby_shto/src/security"
	"encoding/json"
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
	Patch(w http.ResponseWriter, req *http.Request)
	// Post
	Post(w http.ResponseWriter, req *http.Request)
	// Delete ...handlerManager
	Delete(w http.ResponseWriter, req *http.Request)
	// Put
	//Put(w http.ResponseWriter, req *http.Request)
	// Обработчик запросов
	Handle(http.ResponseWriter, *http.Request)
}

// Класс реализующий транспортный уровень
type handlerManager struct {
	Sm security.SessionHandler
}

// Создание инстанса
func CreateInstance(sm *security.SessionManager) Handler {
	return &handlerManager{
		Sm: sm,
	}
}

// Get
func (h *handlerManager) Get(w http.ResponseWriter, req *http.Request) {
	if id := req.URL.Query().Get("id"); id != "" {
		if err := h.fetchOne(id, w); err != nil {
			return
		}
	} else {
		if err := h.fetchList(w); err != nil {
			return
		}
	}
}

func (h *handlerManager) fetchList(w http.ResponseWriter) error {
	var err error
	if err != nil {
		errors.ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError, err)
		return err
	}
	b, _ := json.Marshal(&FetchResult{
		Name: "List books",
		//Total: count,
		Items: nil,
	})
	w.Write([]byte(b))
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *handlerManager) fetchOne(id string, w http.ResponseWriter) error {
	return nil
}

// Post
func (h *handlerManager) Post(w http.ResponseWriter, req *http.Request) {

}

// Patch
func (h *handlerManager) Patch(w http.ResponseWriter, req *http.Request) {

}

// Delete
func (h *handlerManager) Delete(w http.ResponseWriter, req *http.Request) {

}

// Http Handle
func (h *handlerManager) Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.Get(w, req)
		return
	case http.MethodPost:
		h.Post(w, req)
		return
	case http.MethodDelete:
		h.Delete(w, req)
	case http.MethodPatch:
		h.Patch(w, req)
	case http.MethodPut:
		// TODO:: create or load record
		//handler.Put(req)
	}
}
