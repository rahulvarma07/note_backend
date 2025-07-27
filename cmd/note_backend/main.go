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
	"github.com/rahulvarma07/note_backend/internal/http/handlers"
)

func main() {
	var cnf config.Config = *config.MustLoad()

	router := mux.NewRouter()

	router.HandleFunc("/create-a-user", handlers.CreateUser(&cnf.Mail)).Methods("POST")

	server := &http.Server{
		Addr: cnf.HttpServer.BaseUrl,
		Handler: router,
	}

	log.Println("started the server")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func(){
		slog.Info("server has started at", slog.String("port number: ", cnf.HttpServer.Port))
		err := http.ListenAndServe(cnf.HttpServer.Port, router)
		if err != nil{
			log.Fatal("error in starting the server", err)
		}
	}()
	<-stop

	con, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(con)
	if err != nil {
		slog.Error("there is an error in shutting down", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
