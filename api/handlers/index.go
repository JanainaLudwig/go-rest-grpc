package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"grpc-rest/core"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(core.DB.Ping())
	_, _ = fmt.Fprint(w, "Welcome!\n")
}
