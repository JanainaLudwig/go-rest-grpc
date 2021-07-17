package handlers

import (
	"encoding/json"
	"grpc-rest/config"
	"grpc-rest/core"
	"log"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, err error) {
	var response core.Error

	val, ok := (err).(*core.Error)
	if !ok {
		response = *core.NewError(nil, err, http.StatusInternalServerError)
	} else {
		response = *val
	}

	if !config.App.Debug {
		response.File = ""
		response.Line = 0
	}

	SendJsonResponse(w, response, response.Status)
}

func SendJsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header()
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error sending json response", err)
	}
}
