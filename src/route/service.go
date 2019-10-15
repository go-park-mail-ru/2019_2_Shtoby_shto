package route

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
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
	r.Use(AccessCORS, AccessLogMiddleware)
	//apiUserPrefix := utils.Join(apiName, ver, "user")

	e := echo.New()
	apiURL := config.GetInstance().FrontendURL
	e.Use(middleware.Logger(), middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{apiURL},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/user", s.ImageSecurityEcho)
	go e.Start(":8081")
	//modelHandler.InitStaticModel("user")
	//modelHandler := transport.CreateModelHandler()
	r.HandleFunc("/docs/", httpSwagger.WrapHandler)
	r.HandleFunc("/", nil)
	r.HandleFunc("/login", s.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/registration", s.Registration).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/logout", s.CheckSession(s.Logout)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user", s.CheckSession(s.UserSecurity)).Methods(http.MethodGet, http.MethodPut, http.MethodOptions)
	r.HandleFunc("/photo", s.CheckSession(s.ImageSecurity)).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions)
	return r
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func AccessCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SetHeaders(&w)
		if (*r).Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
