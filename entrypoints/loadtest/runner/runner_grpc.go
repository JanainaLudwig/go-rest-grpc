package runner

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"google.golang.org/grpc"
	"grpc-rest/grpc/proto"
	"log"
)

type Grpc struct {
	ctx context.Context
	client proto.StudentsServiceClient
	loadType string
}

func NewRunnerGrpc(host string, loadType string, loads ...Load) *Runner {
	ctx := context.Background()
	conn, e := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	return &Runner{
		ctx:   ctx,
		loads: loads,
		code: "grpc",
		client: &Grpc{
			loadType: loadType,
			client: client,
			ctx: ctx,
		},
	}
}

func (r *Grpc) TestFunc() error {
	if r.loadType == "post" {
		_, err := r.client.CreateStudent(r.ctx, &proto.CreateStudentRequest{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
		})

		return err
	}
	_, err := r.client.GetStudents(r.ctx, &proto.GetStudentsRequest{})

	return err
}
