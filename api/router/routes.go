package router

import (
	"github.com/julienschmidt/httprouter"
	"grpc-rest/api/handlers"
	"net/http"
)

func Routes() http.Handler {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/students", handlers.GetStudents)

	return router
}