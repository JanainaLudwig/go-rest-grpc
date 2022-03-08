package runner

import (
	"context"
	"google.golang.org/grpc"
	"grpc-rest/grpc/proto"
	"log"
)

type Grpc struct {
	ctx context.Context
	client proto.StudentsServiceClient
}

func NewRunnerGrpc(host string, loads ...Load) *Runner {
	ctx := context.Background()
	conn, e := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	return &Runner{
		ctx:   ctx,
		loads: loads,
		client: &Grpc{
			client: client,
			ctx: ctx,
		},
	}
}

func (r *Grpc) TestFunc()  {
	_, err := r.client.GetStudents(r.ctx, &proto.GetStudentsRequest{})
	if err != nil {
		log.Println(err)
	}
}
