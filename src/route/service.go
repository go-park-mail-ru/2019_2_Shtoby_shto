package route

import (
	"2019_2_Shtoby_shto/src/security"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

const (
	apiName = "api"
	ver     = "v1"
)

// swag init

// ShowAccount godoc
// @Summary Show a account
// @Description get user by ID
// @ID get-user-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} model.Error
// @Failure 404 {object} model.Error
// @Failure 500 {object} myError
// @Router /user/{id} [get]

// @title Sample Project API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1

func NewRouterService(s security.Security) *mux.Router {
	r := mux.NewRouter()
	//apiUserPrefix := utils.Join(apiName, ver, "user")
	r.HandleFunc("/docs/", httpSwagger.WrapHandler)
	r.HandleFunc("/", nil)
	r.HandleFunc("/login", s.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/logout", s.Logout).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/registration", s.Registration).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/{id}", s.CheckSession(s.UpdateUserSecurity)).Methods(http.MethodPut, http.MethodOptions)
	return r
}
