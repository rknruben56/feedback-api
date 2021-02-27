package template

import (
	"time"

	"github.com/rknruben56/feedback-api/entity"
)

// Service template usecase
type Service struct {
	repo Repository
}

// NewService creates a new Template Service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateTemplate create a template
func (s *Service) CreateTemplate(class string, content string) (entity.ID, error) {
	t, err := entity.NewTemplate(class, content)
	if err != nil {
		return t.ID, err
	}
	return s.repo.Create(t)
}

// GetTemplate gets a template
func (s *Service) GetTemplate(id entity.ID) (*entity.Template, error) {
	t, err := s.repo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

// ListTemplates list templates
func (s *Service) ListTemplates() ([]*entity.Template, error) {
	templates, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(templates) == 0 {
		return nil, entity.ErrNotFound
	}
	return templates, nil
}

// DeleteTemplate deletes a template
func (s *Service) DeleteTemplate(id entity.ID) error {
	_, err := s.GetTemplate(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateTemplate updates a book
func (s *Service) UpdateTemplate(e *entity.Template) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt.Time = time.Now()
	return s.repo.Update(e)
}
