package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)

	log.Printf("Server started on %s", host)

	http.ListenAndServe(host, r)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Ping")
}
