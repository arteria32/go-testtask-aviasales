package main

import (
	"log"
	"main/handlers"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starts")
	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	router.HandleFunc("/", handlers.HandlerGetHelloWorld)
	router.HandleFunc("/GetAllTickets", handlers.HandlerGetTicketsByInterval)
	router.HandleFunc("/GetBestTicketByType/{type}", handlers.HandlerBestTicketByType)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldnt ListenAndServe()", err)
	}
}
