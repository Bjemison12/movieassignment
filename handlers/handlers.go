package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"movieassignment/entities"
	"movieassignment/repository"
	//"movieassignment/service"
	"net/http"
)

//type MovieService interface {
//	PostMovieHandler(w http.ResponseWriter, r *http.Request)
//	GetAllMovies(w http.ResponseWriter, r *http.Request)
//	GetMovieById(w http.ResponseWriter, r *http.Request)
//	UpdateMovie(w http.ResponseWriter, r *http.Request)
//	DeleteMovie(w http.ResponseWriter, r *http.Request)
//}

type MovieService interface {
	CreateNewMovie(mv entities.Movie) error
	GetAll() (repository.MvStruct, error)
	GetByID(id string) (entities.Movie, error)
	UpdateByID(id string, m entities.Movie) error
	DeleteByID(id string) error
}

type MovieHandler struct {
	Svc MovieService
}

func NewMovieHandler(s MovieService) MovieHandler {
	return MovieHandler{
		Svc: s,
	}
}

func (mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	mMovDB, _ := json.MarshalIndent(mhDB, "", " ")
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(mMovDB)

}

func (mh MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	mvId := mux.Vars(r)
	getId := mvId["id"]

	selectedMovie, err := mh.Svc.GetByID(getId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	mvResponse, _ := json.MarshalIndent(selectedMovie, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(mvResponse)
}

func (mh MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	mvID := mux.Vars(r)
	id := mvID["id"]
	m := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}

	err = mh.Svc.UpdateByID(id, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (mh MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	mvID := mux.Vars(r)
	id := mvID["id"]

	err := mh.Svc.DeleteByID(id)
	if err != nil {
		switch err.Error() {
		case "request not valid":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "delete request failed, movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
		case "Server issue, try again":
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
