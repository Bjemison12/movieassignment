package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"movieassignment/entities"
)

type MvStruct struct {
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepository(filename string) Repo {
	return Repo{
		Filename: filename,
	}
}

func (r Repo) CreateNewMovie(mv entities.Movie) error {
	ms := MvStruct{}

	jsonBytes, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &ms)
	if err != nil {
		return err
	}

	ms.Movies = append(ms.Movies, mv)

	byteSlice, err := json.MarshalIndent(ms, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.Filename, byteSlice, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) GetAll() (MvStruct, error) {
	fn, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	ms := MvStruct{}
	err = json.Unmarshal(fn, &ms)
	if err != nil {
		return ms, err
	}
	return ms, nil
}

func (r Repo) GetByID(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	m := MvStruct{}
	err = json.Unmarshal(file, &m)

	found := entities.Movie{}

	for _, movieId := range m.Movies {
		if movieId.Id == id {
			found = movieId
			return found, nil
		}
	}
	return entities.Movie{}, errors.New("movie not found")
}

func (r Repo) UpdateByID(id string, m entities.Movie) error {

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	mvDb := MvStruct{}
	err = json.Unmarshal(file, &mvDb)

	for i, v := range mvDb.Movies {
		if v.Id == id {
			mvDb.Movies[i] = m
		}
	}
	marshal, err := json.MarshalIndent(mvDb.Movies, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.Filename, marshal, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) DeleteByID(id string) error {
	//instance of slice of movie struct
	m := MvStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &m)
	if err != nil {
		return err
	}

	for index, value := range m.Movies {
		if value.Id == id {
			m.Movies = append(m.Movies[:index], m.Movies[index+1:]...)
			log.Printf("Movie Array: %+v", m.Movies)
		}
	}

	marshal, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.Filename, marshal, 0644)
	if err != nil {
		return err
	}

	return nil
}
