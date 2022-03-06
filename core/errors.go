package core

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"runtime"
)

type Error struct {
	Id interface{} `json:"id,omitempty"`
	Message string `json:"message"`
	Status int `json:"status"`
	File string `json:"file,omitempty"`
	Line int `json:"line,omitempty"`
	Stack string `json:"stack,omitempty"`
	grpcErrorCode codes.Code
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) mapStatus(status int) codes.Code {
	switch status {
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	default:
		return codes.Unknown
	}
}

func (e *Error) ToGrpcError() error {
	return status.Error(e.grpcErrorCode, e.Message)
}

func GrpcError(err error) error {
	if err, ok := err.(*Error); ok {
		return err.ToGrpcError()
	}

	return status.Error(codes.Unknown, err.Error())
}

func WrapError(err error) error {
	_, file, line, _ := runtime.Caller(1)

	if errWrap, ok := err.(*Error); ok {
		errWrap.Stack = fmt.Sprintf("%v:%v : %v", file, line, errWrap.Stack)
	}

	return NewError(0, err, 0)
}

func NotFoundError(id interface{}, message interface{}) *Error {
	return NewError(id, message, http.StatusNotFound)
}

func NewError(id interface{}, message interface{}, status int) *Error {
	_, file, line, _ := runtime.Caller(1)

	if status == 0 {
		status = http.StatusInternalServerError
	}

	e := Error{
		Id: id,
		File: file,
		Line: line,
		Status: status,
	}

	e.grpcErrorCode = e.mapStatus(status)

	switch val := message.(type) {
	case error:
		e.Message = val.Error()
	case string:
		e.Message = val
	}

	return &e
}
