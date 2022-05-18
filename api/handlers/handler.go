package handlers

import (
	"compress/gzip"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"grpc-rest/config"
	"grpc-rest/core"
	"log"
	"net/http"
)

type ResponseCreated struct {
	Id interface{} `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseId struct {
	Id interface{} `json:"id,omitempty"`
}

func Decode(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

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
		response.Stack = ""
	}

	SendJsonResponse(w, response, response.Status)
}

func SendJsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header()
	w.WriteHeader(status)

	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error sending json response", err)
	}
}

func SendJsonResponseGzip(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header()
	w.WriteHeader(status)

	if data == nil {
		return
	}
	writer := gzip.NewWriter(w)
	defer writer.Close()
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Println("Error sending json response", err)
	}
}

func SendProtoResponseGzip(w http.ResponseWriter, data proto.Message, status int) {
	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header()
	w.WriteHeader(status)

	marshal, err := proto.Marshal(data)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	writer := gzip.NewWriter(w)
	defer writer.Close()
	_, err = writer.Write(marshal)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}
}

func SendProtoResponse(w http.ResponseWriter, data proto.Message, status int) {
	if data == nil {
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header()
	w.WriteHeader(status)

	marshal, err := proto.Marshal(data)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}
}
