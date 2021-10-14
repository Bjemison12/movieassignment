package service

import (
	"errors"
	"github.com/google/uuid"
	"movieassignment/entities"
	"movieassignment/repository"
)

type MovieRepository interface {
	CreateNewMovie(mv entities.Movie) error
	GetAll() (repository.MvStruct, error)
	GetByID(id string) (entities.Movie, error)
	UpdateByID(id string, m entities.Movie) error
	DeleteByID(id string) error
}

type Service struct {
	Repo MovieRepository
}

func NewService(r MovieRepository) Service {
	return Service{
		r,
	}
}

func (s Service) CreateNewMovie(mv entities.Movie) error {
	newMV := entities.Movie{}
	mv.Id = uuid.New().String() // this creates a new UUID with the movie when its created.

	if mv.Rating >= 0 && mv.Rating <= 10 {
		err := s.Repo.CreateNewMovie(newMV)
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

func (s Service) GetByID(id string) (entities.Movie, error) {
	m, err := s.Repo.GetByID(id)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (s Service) UpdateByID(id string, m entities.Movie) error {
	if id != m.Id {
		return errors.New("id in body must match url id")
	}
	err := s.Repo.UpdateByID(id, m)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteByID(id string) error {
	err := s.Repo.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil

}
