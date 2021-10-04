package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"movieassignment/entities"
)

type MvStruct struct{
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepository(filename string) Repo{
	return Repo{
		Filename: filename,
	}
}

func (r Repo)CreateNewMovie(mv entities.Movie) error{
	ms := MvStruct{}

	jsonBytes, err := ioutil.ReadFile(r.Filename)
	if err != nil{
		return err
	}
	err = json.Unmarshal(jsonBytes, &ms)
	if err != nil{
		return err
	}

	ms.Movies = append(ms.Movies, mv)

	byteSlice, err:= json.MarshalIndent(ms, "", " ")
	if err != nil{
		return err
	}
	err = ioutil.WriteFile(r.Filename, byteSlice, 0644)
	if err != nil{
		return err
	}

	return nil
}

func (r Repo) GetAll() (MvStruct, error){
	file,err:= ioutil.ReadFile(r.Filename)
	if err != nil{
		fmt.Println(err)
	}
	ms := MvStruct{}
	err =json.Unmarshal(file, &ms)
	if err != nil{
		return ms, err
	}
	return ms,err
}

func (r Repo) GetByID(id string) (entities.Movie, error){
	file,err:= ioutil.ReadFile(r.Filename)
	if err != nil{
		fmt.Println(err)
	}

	m := MvStruct{}
	err =json.Unmarshal(file, &m)

	found := entities.Movie{}
	for _, movie:= range m.Movies{
		if found.Id == id{
			found = movie
			return found, nil
		}
	}
	return found, nil
}