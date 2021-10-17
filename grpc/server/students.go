package server

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-rest/grpc/proto"
	"grpc-rest/models/student"
	"log"
)

type StudentsService struct {
	proto.UnimplementedStudentsServiceServer
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
		resp.Students = append(resp.Students, studentToProto(&u))
	}

	return resp, nil
}

func (s *StudentsService) CreateStudent(ctx context.Context, request *proto.CreateStudentRequest) (*proto.Student, error) {
	std := student.Student{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
	}

	idInserted, err := student.Insert(ctx, &std)
	if err != nil {
		return nil, err
	}

	data, err := student.FetchById(ctx, idInserted)
	if err != nil {
		return nil, err
	}

	return studentToProto(data), nil
}


func studentToProto(s *student.Student) *proto.Student {
	return &proto.Student{
		Id:        int64(s.Id),
		FirstName:  s.FirstName,
		LastName:   s.LastName,
		Identifier: s.Identifier,
		CreatedAt: timestamppb.New(*s.CreatedAt),
		UpdatedAt: timestamppb.New(*s.UpdatedAt),
	}
}
