package main

import (
	"log"
	"movieassignment/handlers"
	"movieassignment/repository"
	"movieassignment/service"
	"net/http"
	"path/filepath"
)

func main() {
	fn := "/Users/biancajemison/Desktop/GoProjects/movieassignment/moviedb.json"

	ext := filepath.Ext(fn)

	if ext != ".json" {
		log.Fatal("File extension is invalid")
	}
	//construct the instance of a repository
	log.Print("This Prints 1st")

	r := repository.NewRepository(fn)
	svc := service.NewService(r)
	hdlr := handlers.NewMovieHandler(svc)
	router := handlers.RouterConfiguration(hdlr)
	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	log.Fatalln(svr.ListenAndServe())
}
