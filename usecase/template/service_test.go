package template

import (
	"testing"
	"time"

	"github.com/rknruben56/feedback-api/entity"
	"github.com/stretchr/testify/assert"
)

func newFixtureTemplate() *entity.Template {
	return &entity.Template{
		Class:     "Class123",
		Content:   "[Student] is doing well",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	e := newFixtureTemplate()

	_, err := s.CreateTemplate(e.Class, e.Content)

	assert.Nil(t, err)
	assert.False(t, e.CreatedAt.IsZero())
}

func Test_List(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	e1 := newFixtureTemplate()
	e2 := newFixtureTemplate()
	_, _ = s.CreateTemplate(e1.Class, e1.Content)
	_, _ = s.CreateTemplate(e2.Class, e2.Content)

	templates, err := s.ListTemplates()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(templates))
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	e := newFixtureTemplate()
	id, err := s.CreateTemplate(e.Class, e.Content)
	assert.Nil(t, err)
	saved, _ := s.GetTemplate(id)
	saved.Class = "Class456"

	assert.Nil(t, s.UpdateTemplate(saved))

	updated, err := s.GetTemplate(id)
	assert.Nil(t, err)
	assert.Equal(t, "Class456", updated.Class)
}

func Test_Delete(t *testing.T) {
	repo := newInmem()
	s := NewService(repo)
	e1 := newFixtureTemplate()
	e2 := newFixtureTemplate()
	e2ID, _ := s.CreateTemplate(e2.Class, e2.Content)

	err := s.DeleteTemplate(e1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = s.DeleteTemplate(e2ID)
	assert.Nil(t, err)
	_, err = s.GetTemplate(e2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
