package main

import (
	"context"
	"encoding/json"
	gproto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/grpc/proto"
	"log"
	"net/http"
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

	getById(ctx, client, 10)
	restGet()
}

func getById(ctx context.Context, client proto.StudentsServiceClient, id int) {
	student, err := client.GetStudentByIdWithSubjects(ctx, &proto.GetStudentByIdRequest{Id: int64(id)})
	if err != nil {
		log.Fatal(err)
	}

	protoSizeBytes := gproto.Size(student)
	log.Println("protoSizeBytes", protoSizeBytes)

	//log.Println(student)
}


func restGet() error {
	ctx := context.Background()
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.App.ServerRest + "/students/10/subjects", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	log.Println(config.App.ServerRest + "/students/1/subjects", "length", res.Header.Get("Content-Length"))
	var students map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&students)
	if err != nil {
		return err
	}

	//log.Println(students)

	return nil
}