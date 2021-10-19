package service

import (
	"errors"
	"github.com/google/uuid"
	"movieassignment/entities"
	"movieassignment/repository"
)

type Repo interface {
	CreateNewMovie(mv entities.Movie) error
	GetAll() (repository.MvStruct, error)
	GetByID(id string) (entities.Movie, error)
	UpdateByID(id string, m entities.Movie) error
	DeleteByID(id string) error
}

type ServRepository struct {
	ServiceRepo Repo
}

func NewService(r Repo) ServRepository {
	return ServRepository{
		r,
	}
}

func (s ServRepository) CreateNewMovie(mv entities.Movie) error {
	mv.Id = uuid.New().String() // this creates a new UUID with the movie when its created.

	if mv.Rating >= 0 && mv.Rating <= 10 {
		err := s.ServiceRepo.CreateNewMovie(mv)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid rating")
}

func (s ServRepository) GetAll() (repository.MvStruct, error) {
	fValue, err := s.ServiceRepo.GetAll()
	if err != nil {
		return fValue, err
	}
	return fValue, nil
}

func (s ServRepository) GetByID(id string) (entities.Movie, error) {
	m, err := s.ServiceRepo.GetByID(id)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (s ServRepository) UpdateByID(id string, m entities.Movie) error {
	if id != m.Id {
		return errors.New("id in body must match url id")
	}
	err := s.ServiceRepo.UpdateByID(id, m)
	if err != nil {
		return err
	}
	return nil
}

func (s ServRepository) DeleteByID(id string) error {
	err := s.ServiceRepo.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil

}
