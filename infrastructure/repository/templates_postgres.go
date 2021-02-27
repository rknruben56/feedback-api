package repository

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rknruben56/feedback-api/entity"
)

// TemplatesPostgres postgres repo
type TemplatesPostgres struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

// NewTemplatesPostgres creates a new repo
func NewTemplatesPostgres(db *sql.DB) *TemplatesPostgres {
	return &TemplatesPostgres{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create a template
func (r *TemplatesPostgres) Create(t *entity.Template) (entity.ID, error) {
	_, err := r.psql.
		Insert("templates").
		Columns("id", "class", "content", "created_at").
		Values(t.ID, t.Class, t.Content, time.Now()).
		RunWith(r.db).
		Exec()

	if err != nil {
		return t.ID, err
	}

	return t.ID, nil
}

// Get template
func (r *TemplatesPostgres) Get(id entity.ID) (*entity.Template, error) {
	rows, err := r.psql.
		Select("*").
		From("templates").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		Query()

	if err != nil {
		return nil, err
	}

	var t entity.Template
	for rows.Next() {
		rows.Scan(&t.ID, &t.Class, &t.Content, &t.CreatedAt, &t.UpdatedAt)
	}

	return &t, nil
}

// Update template
func (r *TemplatesPostgres) Update(t *entity.Template) error {
	t.UpdatedAt.Time = time.Now()
	_, err := r.psql.
		Update("templates").
		SetMap(sq.Eq{"id": t.ID, "class": t.Class, "content": t.Content, "updated_at": t.UpdatedAt}).
		RunWith(r.db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

// List templates
func (r *TemplatesPostgres) List() ([]*entity.Template, error) {
	rows, err := r.psql.
		Select("*").
		From("templates").
		RunWith(r.db).
		Query()

	if err != nil {
		return nil, err
	}

	var templates []*entity.Template
	for rows.Next() {
		var t entity.Template
		err = rows.Scan(&t.ID, &t.Class, &t.Content, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		templates = append(templates, &t)
	}

	return templates, nil
}

// Delete template
func (r *TemplatesPostgres) Delete(id entity.ID) error {
	_, err := r.psql.
		Delete("templates").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		Exec()

	if err != nil {
		return err
	}
	return nil
}
