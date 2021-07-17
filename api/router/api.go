package router

import (
	"grpc-rest/config"
	"log"
	"net/http"
)

func StartApi() {
	for  {
		log.Println(http.ListenAndServe(":" + config.App.ApiPort, Routes()))
	}
}
