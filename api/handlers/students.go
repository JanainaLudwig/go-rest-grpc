package handlers

import (
	"github.com/julienschmidt/httprouter"
	"grpc-rest/domain"
	"grpc-rest/repositories/student"
	"grpc-rest/repositories/student_subject"
	"net/http"
	"strconv"
)

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	all, err := student.FetchAll(r.Context())
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, all, http.StatusOK)
}

func GetStudentById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id_student"))
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	all, err := student.FetchById(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, all, http.StatusOK)
}

func GetStudentSubjectsById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id_student"))
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	std, err := student.FetchById(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	subjects, err := student_subject.FetchByStudentSubjectId(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, struct {
		domain.Student
		Subjects []domain.StudentSubjectWithSubject `json:"subjects"`
	}{
		Student:  *std,
		Subjects: subjects,
	}, http.StatusOK)
}

func CreateStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := domain.Student{}

	err := Decode(r, &s)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	id, err := student.Create(r.Context(), &s)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, ResponseId{
		Id:      id,
	}, http.StatusCreated)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s := domain.Student{}

	err := Decode(r, &s)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	s.Id, err = strconv.Atoi(p.ByName("id_student"))
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	err = student.Update(r.Context(), &s)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, nil, http.StatusOK)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))

	err = student.Delete(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, nil, http.StatusOK)
}
