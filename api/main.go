package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/api/handler"
	"github.com/rknruben56/feedback-api/api/middleware"
	"github.com/rknruben56/feedback-api/usecase/template"
)

func main() {
	templateRepo := template.NewInmem()
	templateService := template.NewService(templateRepo)

	r := mux.NewRouter()
	handler.MakeTemplateHandlers(r, templateService)
	http.Handle("/", r)
	r.Use(middleware.Cors)

	fmt.Println("listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
