package main

import (
	"2019_2_Shtoby_shto/chat/app"
	"2019_2_Shtoby_shto/session_service/session"
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/initDB"
	"2019_2_Shtoby_shto/src/security"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(echoLog.DEBUG)
	if err := config.InitConfig(); err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}

	conf := config.GetInstance()

	sessService, err := ConnectGRPC(conf.SecurityURL, "security_service")

	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	securityClient := session.NewSecurityClient(sessService)

	httpAddr := ":" + strconv.Itoa(conf.Port)
	log.Info("API Url:", httpAddr, securityClient)

	dbService := initDB.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		e.Logger.Error(err)
		os.Exit(3)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	InitServices(e, dm, conf, securityClient)
	newServer(e, httpAddr)

	e.GET("/ws", func(ctx echo.Context) error {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return err
		}
		conn.NextReader()
		return nil
	})

	server := app.NewServer("/ws")
	go server.Listen()

	//e.GET("/chat", nil)
	http.Handle("/", http.FileServer(http.Dir("localhost")))
	log.Fatal(http.ListenAndServe(":8082", nil))

	if err := e.Start("localhost:8081"); err != http.ErrServerClosed {
		e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func newServer(e *echo.Echo, httpAddr string) {
	e.Logger.Info("serving on", httpAddr)

	apiURL := config.GetInstance().FrontendURL
	e.Use(
		middleware.Logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{apiURL},
			AllowCredentials: true,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
			ExposeHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		}))

	e.Server = &http.Server{
		Addr:           httpAddr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func InitServices(e *echo.Echo, db database.IDataManager, conf *config.Config, sessService session.SecurityClient) {
	sessionService := security.NewSessionManager(&sessService)
	securityService := security.CreateInstance(sessionService)
	e.Use(securityService.CheckSession)
}

func ConnectGRPC(addr string, name string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	if err != nil {
		log.Error("Can't connect to security service:", err)
		return nil, err
	}
	return conn, nil
}
