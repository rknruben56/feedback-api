package template

import "github.com/rknruben56/feedback-api/entity"

// Reader...
type Reader interface {
	Get(id entity.ID) (*entity.Template, error)
	List() ([]*entity.Template, error)
	Search(query string) ([]*entity.Template, error)
}

// Writer...
type Writer interface {
	Create(e *entity.Template) (entity.ID, error)
	Update(e *entity.Template) error
	Delete(id entity.ID) error
}

// Repository...
type Repository interface {
	Reader
	Writer
}

// UseCase...
type UseCase interface {
	GetTemplate(id entity.ID) (*entity.Template, error)
	ListTemplates() ([]*entity.Template, error)
	CreateTemplate(class string, content string) (entity.ID, error)
	UpdateTemplate(e *entity.Template) error
	DeleteTemplate(id entity.ID) error
	SearchTemplates(query string) ([]*entity.Template, error)
}
