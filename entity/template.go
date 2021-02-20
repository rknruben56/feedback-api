package entity

import "time"

type Template struct {
	ID        ID
	Class     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTemplate creates a new template
func NewTemplate(class string, content string) (*Template, error) {
	t := &Template{
		ID:        NewID(),
		Class:     class,
		Content:   content,
		CreatedAt: time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

// Validate validates a template
func (t *Template) Validate() error {
	if t.Class == "" || t.Content == "" {
		return ErrInvalidEntity
	}
	return nil
}
