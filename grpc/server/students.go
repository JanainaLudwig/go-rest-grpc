package server

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-rest/core"
	"grpc-rest/domain"
	"grpc-rest/grpc"
	"grpc-rest/grpc/proto"
	"grpc-rest/models/student"
	"grpc-rest/models/student_subject"
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
		return nil, core.GrpcError(err)
	}

	resp := &proto.GetStudentsResponse{}
	for _, u := range students {
		resp.Students = append(resp.Students, s.studentToProto(&u))
	}

	return resp, nil
}

func (s *StudentsService) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	std := domain.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	idInserted, err := student.Create(ctx, &std)
	if err != nil {
		return nil, core.GrpcError(err)
	}

	return &proto.CreateStudentResponse{
		Id: int64(idInserted),
	}, nil
}

func (s *StudentsService) GetStudentById(ctx context.Context, req *proto.GetStudentByIdRequest) (*proto.GetStudentByIdResponse, error) {
	fetchById, err := student.FetchById(ctx, int(req.Id))
	if err != nil {
		return nil, core.GrpcError(err)
	}

	return &proto.GetStudentByIdResponse{
		Student: s.studentToProto(fetchById),
	}, nil
}

func (s *StudentsService) UpdateStudentById(ctx context.Context, req *proto.UpdateStudentByIdRequest) (*proto.UpdateStudentByIdResponse, error) {
	err := student.Update(ctx, &domain.Student{
		Id:        int(req.Id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, core.GrpcError(err)
	}

	return &proto.UpdateStudentByIdResponse{}, nil
}

func (s *StudentsService) DeleteStudentById(ctx context.Context, req *proto.DeleteStudentByIdRequest) (*proto.DeleteStudentByIdResponse, error) {
	err := student.Delete(ctx, int(req.Id))

	if err != nil {
		return nil, core.GrpcError(err)
	}

	return &proto.DeleteStudentByIdResponse{}, nil
}

func (s *StudentsService) GetStudentByIdWithSubjects(ctx context.Context, req *proto.GetStudentByIdRequest) (*proto.GetStudentByIdWithSubjectsResponse, error) {
	id := int(req.Id)
	std, err := student.FetchById(ctx, id)
	if err != nil {
		return nil, core.GrpcError(err)
	}

	subjects, err := student_subject.FetchByStudentSubjectId(ctx, id)
	if err != nil {
		return nil, core.GrpcError(err)
	}

	var protoSubjects []*proto.StudentSubject
	for _, subject := range subjects {
		protoSubjects = append(protoSubjects, s.studentSubjectToProto(&subject))
	}

	return &proto.GetStudentByIdWithSubjectsResponse{
		Student:  s.studentToProto(std),
		Subjects: protoSubjects,
	}, nil
}

func (s *StudentsService) studentToProto(std *domain.Student) *proto.Student {
	return &proto.Student{
		Id:         int64(std.Id),
		FirstName:  std.FirstName,
		LastName:   std.LastName,
		Identifier: std.Identifier,
		CreatedAt:  timestamppb.New(*std.CreatedAt),
		UpdatedAt:  timestamppb.New(*std.UpdatedAt),
	}
}

func (s *StudentsService) studentSubjectToProto(std *domain.StudentSubjectWithSubject) *proto.StudentSubject {
	return &proto.StudentSubject{
		Id:        int64(std.Id),
		IdSubject: int64(std.IdSubject),
		Frequency: float32(std.Frequency),
		Status:    std.Status,
		CreatedAt: timestamppb.New(*std.CreatedAt),
		UpdatedAt: timestamppb.New(*std.UpdatedAt),
		Name:      std.Name,
	}
}

func (s *StudentsService) protoToStudent(std *proto.Student) *domain.Student {
	return &domain.Student{
		Id:         int(std.Id),
		FirstName:  std.FirstName,
		LastName:   std.LastName,
		Identifier: std.Identifier,
		ModelDate: domain.ModelDate{
			CreatedAt: grpc.PrototimeToTime(std.CreatedAt),
			UpdatedAt: grpc.PrototimeToTime(std.UpdatedAt),
		},
	}
}
