package handlers

import (
	"github.com/gorilla/mux"
)

func RouterConfiguration(handler MovieHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	router.HandleFunc("/movie", handler.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", handler.GetMovieByID).Methods("GET")
	//router.HandleFunc("/movie/{id}", handler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/delete/{id}", handler.DeleteMovie).Methods("DELETE")

	return router
}
