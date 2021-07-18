package server

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-rest/grpc/proto"
	"grpc-rest/models/student"
	"log"
)

type StudentsService struct {
}

func NewStudentsServiceController() *StudentsService {
	return &StudentsService{}
}

func (s *StudentsService) GetStudents(ctx context.Context, req *proto.GetStudentsRequest) (*proto.GetStudentsResponse, error) {
	students, err := student.FetchAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &proto.GetStudentsResponse{}
	for _, u := range students {
		resp.Students = append(resp.Students, marshalStudent(&u))
	}

	return resp, nil
}

func marshalStudent(s *student.Student) *proto.Student {
	return &proto.Student{
		Id:        int64(s.Id),
		FirstName:  s.FirstName,
		LastName:   s.LastName,
		Identifier: s.Identifier,
		CreatedAt: timestamppb.New(*s.CreatedAt),
		UpdatedAt: timestamppb.New(*s.UpdatedAt),
	}
}
