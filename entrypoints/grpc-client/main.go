package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/grpc/proto"
	"log"
	gproto "github.com/golang/protobuf/proto"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")
	core.StartApp(true)
	ctx := context.Background()

	conn, e := grpc.DialContext(ctx, config.App.ServerGrpc, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	getById(ctx, client, 1)
}

func listAll(ctx context.Context, client proto.StudentsServiceClient) {
	students, e := client.GetStudents(ctx, &proto.GetStudentsRequest{})
	if e != nil {
		log.Fatal(e)
		return
	}

	log.Println("password")
	log.Println("password")
	log.Println("password")
	log.Println("password")
	log.Println("password")
	log.Println("password")
	log.Println("password2")
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
	student, err := client.GetStudentByIdWithSubjects(ctx, &proto.GetStudentByIdRequest{Id: int64(id)})
	if err != nil {
		log.Fatal(err)
	}

	protoSizeBytes := gproto.Size(student)
	log.Println("protoSizeBytes", protoSizeBytes)

	log.Println(student)
}

