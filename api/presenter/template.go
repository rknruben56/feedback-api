package presenter

import (
	"time"

	"github.com/rknruben56/feedback-api/entity"
)

// Template data
type Template struct {
	ID        entity.ID `json:"id"`
	Class     string    `json:"class"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
