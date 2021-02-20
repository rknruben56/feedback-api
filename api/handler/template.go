package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/api/presenter"
	"github.com/rknruben56/feedback-api/entity"
	"github.com/rknruben56/feedback-api/usecase/template"
)

func listTemplates(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading templates"
		data, err := service.ListTemplates()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Template
		for _, d := range data {
			toJ = append(toJ, &presenter.Template{
				ID:        d.ID,
				Class:     d.Class,
				Content:   d.Content,
				CreatedAt: d.CreatedAt,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func createTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding template"
		var input struct {
			Class   string `json:"class"`
			Content string `json:"content"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateTemplate(input.Class, input.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Template{
			ID:      id,
			Class:   input.Class,
			Content: input.Content,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func MakeTemplateHandlers(r *mux.Router, service template.UseCase) {
	r.Handle("/v1/template", listTemplates(service)).Methods("GET", "OPTIONS").Name("listTemplates")
	r.Handle("/v1/template", createTemplate(service)).Methods("POST", "OPTIONS").Name("createTemplate")
}
