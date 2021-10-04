package handlers

import (
	"github.com/gorilla/mux"
)

func RouterConfiguration(handler  MovieHandler) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	router.HandleFunc("/movie", handler.GetMovieByID).Methods("GET")
	router.HandleFunc("/movie/{Id}", handler.GetMovieByID).Methods("GET")

	return router
}
