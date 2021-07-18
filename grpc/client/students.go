package client

import (
	"context"
	"google.golang.org/grpc"
	"grpc-rest/grpc/proto"
)

type StudentsClient struct {
}

func NewStudentsClient () *StudentsClient {
	return &StudentsClient{}
}

func (s *StudentsClient) GetStudents(ctx context.Context, in *proto.GetStudentsRequest, opts ...grpc.CallOption) (*proto.GetStudentsResponse, error) {
	panic("implement me")
}

//func (s *StudentsClient) GetStudents(ctx context.Context, in *proto.GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error)  {
//
//}
