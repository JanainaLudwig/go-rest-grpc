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
	core.StartApp(true)
	ctx := context.Background()

	serverAddress := "localhost:" + config.App.GrpcPort
	conn, e := grpc.DialContext(ctx, serverAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	listAll(ctx, client)
	id := createSample(ctx, client)
	getById(ctx, client, id)
	getById(ctx, client, 1550)
}

func listAll(ctx context.Context, client proto.StudentsServiceClient) {
	students, e := client.GetStudents(ctx, &proto.GetStudentsRequest{})
	if e != nil {
		log.Fatal(e)
		return
	}

	data, err := json.MarshalIndent(students.GetStudents(), "", " ")
	log.Println(string(data), err)
}

func createSample(ctx context.Context, client proto.StudentsServiceClient) int {
	student, err := client.CreateStudent(ctx, &proto.CreateStudentRequest{
		FirstName: "Janaina",
		LastName:  "Ludwig",
	})
	if err != nil {
		log.Fatal(err)
	}

	return int(student.Id)
}


func getById(ctx context.Context, client proto.StudentsServiceClient, id int) {
	student, err := client.GetStudentById(ctx, &proto.GetStudentByIdRequest{Id: int64(id)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(student)
}

