package route

import (
	"2019_2_Shtoby_shto/src/security"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	apiName = "api"
	ver     = "v1"
)

func NewRouterService(s security.Security) *mux.Router {
	r := mux.NewRouter()
	//apiUserPrefix := utils.Join(apiName, ver, "user")
	r.HandleFunc("/", nil)
	r.HandleFunc("/login", s.Login).Methods(http.MethodPost)
	return r
}
