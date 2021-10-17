package core

import (
	"fmt"
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
}

func (e *Error) Error() string {
	return e.Message
}

func WrapError(err error) error {
	_, file, line, _ := runtime.Caller(1)

	if errWrap, ok := err.(*Error); ok {
		errWrap.Stack = fmt.Sprintf("%v:%v : %v", file, line, errWrap.Stack)
	}

	return NewError(0, err, 0)
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

	switch val := message.(type) {
	case error:
		e.Message = val.Error()
	case string:
		e.Message = val
	}

	return &e
}
