package handlers

import (
	"github.com/julienschmidt/httprouter"
	"grpc-rest/models/student"
	"net/http"
)

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	all, err := student.FetchAll(r.Context())
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	SendJsonResponse(w, all, http.StatusOK)
}
