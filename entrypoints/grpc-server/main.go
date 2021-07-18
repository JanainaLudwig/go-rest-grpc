package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/grpc/proto"
	services "grpc-rest/grpc/server"
	"net"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")
	core.StartApp()

	server := grpc.NewServer()

	proto.RegisterStudentsServiceServer(server, services.NewStudentsServiceController())

	reflection.Register(server)

	con, err := net.Listen("tcp", ":" + config.App.GrpcPort)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}

