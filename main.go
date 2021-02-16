package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/handlers"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/templates", handlers.GetTemplates)
	log.Println("Listening on port 8080...")
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}
