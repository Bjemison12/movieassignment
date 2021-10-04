package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"movieassignment/entities"
	"movieassignment/service"
	"net/http"
)

type MovieHandler struct {
	Svc service.Service
}

func NewMovieHandler(s service.Service) MovieHandler  {
	return MovieHandler{
		Svc: s,
	}
}

func(mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request){
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil{
		fmt.Println(err)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil{
		switch err.Error() {
		case "movie already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (mh MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	mhDB, err := mh.Svc.GetAll()
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mMovDB, _ := json.MarshalIndent(mhDB, "", " ")
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(mMovDB)

}

func (mh MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request){
	mvId := mux.Vars(r)
	getId := mvId["Id"]
	selectedMovie,err := mh.Svc.GetById(getId)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	mvResponse, _ := json.MarshalIndent(selectedMovie, "", " ")
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(mvResponse)
}