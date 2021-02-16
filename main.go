package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/templates", handlers.GetTemplates)
	log.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
