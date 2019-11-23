package main

import (
	"2019_2_Shtoby_shto/session_service/config"
	"2019_2_Shtoby_shto/session_service/session"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		println(err)
		os.Exit(1)
	}
	conf := config.GetInstance()
	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sHandler := session.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)

	//NewSecurityClient()
	server := grpc.NewServer()
	session.RegisterSecurityServer(server, sHandler)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
