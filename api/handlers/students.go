package handlers

import (
	"github.com/julienschmidt/httprouter"
	"grpc-rest/models/student"
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

func CreateStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := student.Student{}

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

	SendJsonResponse(w, ResponseCreated{
		Id:      id,
		Message: "Student created",
	}, http.StatusCreated)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s := student.Student{}

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

	SendJsonResponse(w, ResponseCreated{
		Id:      s.Id,
		Message: "Student updated",
	}, http.StatusCreated)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))

	err = student.Delete(r.Context(), id)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, ResponseCreated{
		Id:      id,
		Message: "Student deleted",
	}, http.StatusCreated)
}
