package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaswdr/faker"
	"github.com/rknruben56/feedback-api/models"
)

func GetTemplates(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buildTemplates())
}

func buildTemplates() []models.Template {
	var templates []models.Template
	faker := faker.New()
	for i := 0; i <= 10; i++ {
		date := fmt.Sprintf("%s %d, %d", faker.Time().MonthName(), faker.Time().DayOfMonth(), faker.Time().Year())
		template := models.Template{
			ID:          i,
			Class:       faker.UUID().V4(),
			Content:     faker.Lorem().Text(250),
			DateCreated: date,
		}

		templates = append(templates, template)
	}

	return templates
}
