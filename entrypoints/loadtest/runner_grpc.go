package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-rest/grpc/proto"
	"log"
)

type RunnerGrpc struct {
	ctx context.Context
	client proto.StudentsServiceClient
	loads []Load
}

func NewRunnerGrpc(host string) *RunnerGrpc {
	ctx := context.Background()
	conn, e := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}

	client := proto.NewStudentsServiceClient(conn)

	return &RunnerGrpc{
		ctx: ctx,
		client: client,
	}
}

func (r *RunnerGrpc) AddLoad(load Load) {
	r.loads = append(r.loads, load)
}

