package main

import (
	"2019_2_Shtoby_shto/file_service/config"
	"2019_2_Shtoby_shto/file_service/file"
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

	fHandler := file.CreateFileLoaderInstance(conf.StorageRegion, conf.StorageEndpoint, conf.StorageBucket)

	server := grpc.NewServer()
	file.RegisterIFileLoaderManagerServer(server, fHandler)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
