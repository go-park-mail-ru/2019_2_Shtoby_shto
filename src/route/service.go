package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	apiName = "api"
	ver     = "v1"
)

func NewRouterService() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(apiName+"/"+ver+"/user", nil)
	http.Handle("/", r)
	return r
}
