package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rahulvarma07/note_backend/internal/config"
	"github.com/rahulvarma07/note_backend/internal/http/database"
	"github.com/rahulvarma07/note_backend/internal/http/handlers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	var cnf config.Config = *config.MustLoad()
	router := mux.NewRouter()

	var MongoClient *mongo.Client = database.MustGetMongoClient()
	var UserAuthMongoCollection *mongo.Collection = database.CreateMongCollection(MongoClient, "userCollection", "credentials")
	var UserNoteMongoCollection *mongo.Collection = database.CreateMongCollection(MongoClient, "notesCollection", "userNotes")

	router.HandleFunc("/note/signup", handlers.UserVerification(&cnf.Mail, UserAuthMongoCollection)).Methods("POST")
	router.HandleFunc("/note/mail-verification", handlers.SignUpUser(UserAuthMongoCollection)).Methods("GET")
	router.HandleFunc("/note/login", handlers.LoginTheUser(UserAuthMongoCollection)).Methods("POST")
	router.HandleFunc("/note/add-note", handlers.AddNotes(UserNoteMongoCollection)).Methods("POST")

	server := &http.Server{
		Addr:    cnf.HttpServer.BaseUrl,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("Server has started", slog.String("port", cnf.HttpServer.Port))
		err := http.ListenAndServe(cnf.HttpServer.Port, router)
		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()
	<-stop

	con, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(con)
	if err != nil {
		slog.Error("there is an error in shutting down", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
