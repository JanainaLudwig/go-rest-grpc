package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/grpc/proto"
	"log"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")
	core.StartApp()

	ctx := context.Background()
	serverAddress := "localhost:" + config.App.GrpcPort
	conn, e := grpc.DialContext(ctx, serverAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)
	students, e := client.GetStudents(ctx, &proto.GetStudentsRequest{})
	if e != nil {
		log.Println(e)
		return
	}

	//log.Println(proto2.Size(students), "bytes")
	log.Println(students.GetStudents())
}

