package router

import (
	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
	"grpc-rest/api/handlers"
	"grpc-rest/config"
	"net/http"
)

func Routes() http.Handler {
	router := nrhttprouter.New(config.App.NewRelicApp)

	router.GET("/", handlers.Index)
	router.GET("/students", handlers.GetStudents)
	router.POST("/students", handlers.CreateStudent)
	router.GET("/students/:id_student/subjects", handlers.GetStudentSubjectsById)
	router.PUT("/students/:id_student", handlers.UpdateStudent)
	router.DELETE("/students/:id", handlers.DeleteStudent)

	router.GET("/students-rest", handlers.GetStudentsRest)
	router.GET("/students/:id_student", handlers.GetStudentById)
	router.GET("/students-proto/:id_student", handlers.GetStudentByIdProto)
	router.GET("/students-proto", handlers.GetStudentsProto)

	return router
}
