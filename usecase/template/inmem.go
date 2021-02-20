package template

import "github.com/rknruben56/feedback-api/entity"

// inmem in memory repository
type inmem struct {
	m map[entity.ID]*entity.Template
}

// newInmem creates a new repository
// TODO: make private
func NewInmem() *inmem {
	var m = map[entity.ID]*entity.Template{}
	return &inmem{
		m: m,
	}
}

// Create template
func (r *inmem) Create(e *entity.Template) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get template
func (r *inmem) Get(id entity.ID) (*entity.Template, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update template
func (r *inmem) Update(e *entity.Template) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List templates
func (r *inmem) List() ([]*entity.Template, error) {
	var t []*entity.Template
	for _, i := range r.m {
		t = append(t, i)
	}
	return t, nil
}

// Delete template
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
