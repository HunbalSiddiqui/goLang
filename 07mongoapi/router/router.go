package router

import (
	"github.com/gorilla/mux"
	"github.com/hunbalsiddiqui/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteMovies).Methods("DELETE")
	return router
}
