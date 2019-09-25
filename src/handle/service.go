package transport

import (
	"2019_2_Shtoby_shto/src/errors"
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

type Router interface {
	// Get
	Get(w http.ResponseWriter, req *http.Request)
	// Patch
	Patch(w http.ResponseWriter, req *http.Request)
	// Post
	Post(w http.ResponseWriter, req *http.Request)
	// Delete ...Handler
	Delete(w http.ResponseWriter, req *http.Request)
	// Put
	//Put(w http.ResponseWriter, req *http.Request)
	// Обработчик запросов
	Handle(http.ResponseWriter, *http.Request)
}

// Класс реализующий транспортный уровень
type HttpHandler struct {
	//Sm    *security.SessionManager
}

// Создание инстанса
func CreateInstance() *HttpHandler {
	return &HttpHandler{
		//Sm:    sm,
	}
}

// Get
func (s *HttpHandler) Get(w http.ResponseWriter, req *http.Request) {
	if id := req.URL.Query().Get("id"); id != "" {
		if err := s.fetchOneBook(id, w); err != nil {
			return
		}
	} else {
		if err := s.fetchList(w); err != nil {
			return
		}
	}
}

func (s *HttpHandler) fetchList(w http.ResponseWriter) error {
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

func (s *HttpHandler) fetchOneBook(id string, w http.ResponseWriter) error {
	return nil
}

// Post
func (s *HttpHandler) Post(w http.ResponseWriter, req *http.Request) {

}

// Patch
func (s *HttpHandler) Patch(w http.ResponseWriter, req *http.Request) {

}

// Delete
func (s *HttpHandler) Delete(w http.ResponseWriter, req *http.Request) {

}

// Http Handle
func (s *HttpHandler) Handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch req.Method {
	case http.MethodGet:
		s.Get(w, req)
		return
	case http.MethodPost:
		s.Post(w, req)
		return
	case http.MethodDelete:
		s.Delete(w, req)
	case http.MethodPatch:
		s.Patch(w, req)
	case http.MethodPut:
		// TODO:: create or load record
		//handler.Put(req)
	}
}
