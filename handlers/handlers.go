package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	ticketsfinder "main/services"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerGetHelloWorld(wr http.ResponseWriter,
	req *http.Request) {
	fmt.Fprintf(wr, "Hello, World\n")
	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}

func HandlerGetTicketsByInterval(wr http.ResponseWriter,
	req *http.Request) {
	values := req.URL.Query()
	airportFrom := values.Get("from")
	airportTo := values.Get("to")
	log.Println(airportFrom)
	log.Println(airportTo)
	if len(airportFrom) == 0 || len(airportTo) == 0 {
		http.Error(wr, fmt.Sprintf("Not found params"), http.StatusNotFound)
		return
	}

}

func HandlerBestTicketByType(wr http.ResponseWriter,
	req *http.Request) {
	values := req.URL.Query()
	airportFrom := values.Get("from")
	airportTo := values.Get("to")
	if len(airportFrom) == 0 || len(airportTo) == 0 {
		http.Error(wr, fmt.Sprintf("Not found params"), http.StatusNotFound)
		return
	}
	vars := mux.Vars(req)
	typeSearch := vars["type"]
	if len(typeSearch) == 0 {
		http.Error(wr, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	res, err := ticketsfinder.GetTicketsInfoByIntervalAndType(airportFrom, airportTo, typeSearch)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("res", res)
	jsonResponse, jsonError := json.Marshal(res)
	if jsonError != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonResponse)
}
