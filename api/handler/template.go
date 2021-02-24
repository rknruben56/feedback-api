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
			respondWithInternalServerError(w, errorMessage)
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
			respondWithInternalServerError(w, errorMessage)
		}
	})
}

func getTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading template"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		data, err := service.GetTemplate(id)
		if err != nil && err != entity.ErrNotFound {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Template{
			ID:        data.ID,
			Class:     data.Class,
			Content:   data.Content,
			CreatedAt: data.CreatedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			respondWithInternalServerError(w, errorMessage)
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
			respondWithInternalServerError(w, errorMessage)
			return
		}

		id, err := service.CreateTemplate(input.Class, input.Content)
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		toJ := &presenter.Template{
			ID:      id,
			Class:   input.Class,
			Content: input.Content,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			respondWithInternalServerError(w, errorMessage)
		}
	})
}

func updateTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating template"
		var input struct {
			ID      string `json:"id"`
			Class   string `json:"class"`
			Content string `json:"content"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		id, err := entity.StringToID(input.ID)
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		data, err := service.GetTemplate(id)
		if err != nil && err != entity.ErrNotFound {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		data.Class = input.Class
		data.Content = input.Content
		err = service.UpdateTemplate(data)
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
	})
}

func deleteTemplate(service template.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing template"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
		err = service.DeleteTemplate(id)
		if err != nil {
			respondWithInternalServerError(w, errorMessage)
			return
		}
	})
}

func respondWithInternalServerError(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errorMessage))
}

func MakeTemplateHandlers(r *mux.Router, service template.UseCase) {
	r.Handle("/v1/templates", listTemplates(service)).
		Methods("GET", "OPTIONS").
		Name("listTemplates")
	r.Handle("/v1/templates/{id}", getTemplate(service)).
		Methods("GET", "OPTIONS").
		Name("getTemplate")
	r.Handle("/v1/templates", createTemplate(service)).
		Methods("POST", "OPTIONS").
		Name("createTemplate")
	r.Handle("/v1/templates", updateTemplate(service)).
		Methods("PUT", "OPTIONS").
		Name("updateTemplate")
	r.Handle("/v1/templates/{id}", deleteTemplate(service)).
		Methods("DELETE", "OPTIONS").
		Name("deleteTemplate")
}
