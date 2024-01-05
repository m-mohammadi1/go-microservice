package memory

import (
	"context"
	"sync"

	"movieexample.com/metadata/internal/repository"
	model "movieexample.com/metadata/pkg"
)

// Repository defines a memory movie metadata
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{},
	}
}

// Get retrieve movie metadata for movie id
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put add movies metadata for given movie id
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.RUnlock()

	r.data[id] = metadata
	return nil
}
