package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahulvarma07/note_backend/internal/config"
	"github.com/rahulvarma07/note_backend/internal/http/handlers"
)

func main() {
	var cnf config.Config = *config.MustLoad()

	router := mux.NewRouter()

	router.HandleFunc("/create-a-user", handlers.CreateUser(&cnf.Mail)).Methods("POST")

	log.Println("started the server")
	err := http.ListenAndServe(":8082", router)
	if err != nil{
		log.Fatal("unable to start the server")
	}
}
