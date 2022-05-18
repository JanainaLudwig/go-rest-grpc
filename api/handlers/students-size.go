package handlers

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-rest/domain"
	"grpc-rest/grpc/proto"
	"grpc-rest/repositories/student"
	"net/http"
	"strconv"
)

const (
	qtdStudentsList = 2000
)


func sendGzip(r *http.Request) bool {
	return r.Header.Get("Content-Encoding") == "gzip"
}


func GetStudentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id_student"))
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	_, std, err := getSampleStudent(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}


	if sendGzip(r) {
		SendJsonResponseGzip(w, std, http.StatusOK)
		return
	}

	SendJsonResponse(w, std, http.StatusOK)
}

func GetStudentsProto(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stds []*proto.Student
	studentProto, _, err := getSampleStudent(r.Context(), 1)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}
	for i := 0; i < qtdStudentsList; i++ {
		stds = append(stds, studentProto)
	}

	if sendGzip(r) {
		SendProtoResponseGzip(w, &proto.GetStudentsResponse{Students: stds}, http.StatusOK)
		return
	}

	SendProtoResponse(w, &proto.GetStudentsResponse{Students: stds}, http.StatusOK)
}

func GetStudentsRest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stds []domain.Student
	_, std, err := getSampleStudent(r.Context(), 1)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}
	for i := 0; i < qtdStudentsList; i++ {
		stds = append(stds, *std)
	}

	if sendGzip(r) {
		SendJsonResponseGzip(w, stds, http.StatusOK)
		return
	}

	SendJsonResponse(w, stds, http.StatusOK)
}

func GetStudentByIdProto(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id_student"))
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	studentProto, _, err := getSampleStudent(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}


	if sendGzip(r) {
		SendProtoResponseGzip(w, studentProto, http.StatusOK)
		return
	}

	SendProtoResponse(w, studentProto, http.StatusOK)
}

func getSampleStudent(ctx context.Context, id int) (*proto.Student, *domain.Student, error) {
	std, err := student.FetchById(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	return &proto.Student{
		Id:         int64(std.Id),
		FirstName:  std.FirstName,
		LastName:   std.LastName,
		Identifier: std.Identifier,
		CreatedAt:  timestamppb.New(*std.CreatedAt),
		UpdatedAt:  timestamppb.New(*std.UpdatedAt),
	}, std, nil
}