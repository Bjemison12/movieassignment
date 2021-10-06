package service

import (
	"errors"
	"github.com/google/uuid"
	"movieassignment/entities"
	"movieassignment/repository"
)

type InvalidRatingError error

type Service struct {
	Repo repository.Repo
}

func NewService(r repository.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateNewMovie(mv entities.Movie) error {

	mv.Id = uuid.New().String()

	if mv.Rating >= 0 && mv.Rating <= 10 {
		err := s.Repo.CreateNewMovie(mv)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid rating")
}

func (s Service) GetAll() (repository.MvStruct, error) {
	fValue, err := s.Repo.GetAll()
	if err != nil {
		return fValue, err
	}
	return fValue, nil
}

func (s Service) GetById(id string) (entities.Movie, error) {
	m, err := s.Repo.GetByID(id)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (s Service) DeleteMovieByID(id string) error {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil

}
