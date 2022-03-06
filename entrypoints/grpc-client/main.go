package main

import (
	"context"
	"encoding/json"
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
		log.Fatal(e)
		return
	}

	data, err := json.MarshalIndent(students.GetStudents(), "", " ")
	log.Println(string(data), err)
}

