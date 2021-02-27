package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/api/handler"
	"github.com/rknruben56/feedback-api/api/middleware"
	"github.com/rknruben56/feedback-api/config"
	"github.com/rknruben56/feedback-api/infrastructure/repository"
	"github.com/rknruben56/feedback-api/usecase/template"

	_ "github.com/lib/pq"
)

func main() {
	dbInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
	)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	templateRepo := repository.NewTemplatesPostgres(db)
	templateService := template.NewService(templateRepo)

	r := mux.NewRouter()
	handler.MakeTemplateHandlers(r, templateService)
	http.Handle("/", r)
	r.Use(middleware.Cors)

	fmt.Println("listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
