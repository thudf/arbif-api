package router

import (
	"arbif/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/leads", controllers.CreateLead).Methods(http.MethodPost)
	return router
}
