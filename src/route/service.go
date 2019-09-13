package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func newService(host string) *mux.Router {
	r := mux.NewRouter()
	r.Host(host)
	r.HandleFunc("/", nil)
	r.HandleFunc("/products", nil)
	r.HandleFunc("/articles", nil)
	http.Handle("/", r)
	return r
}
