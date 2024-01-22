package metadata

import (
	"context"
	"errors"

	"movieexample.com/metadata/internal/repository"
	"movieexample.com/metadata/pkg/model"
)

// ErrNotFound when cannot find record
var ErrNotFound = errors.New("not found")

type metadataRepositoryInterface interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller
type Controller struct {
	repo metadataRepositoryInterface
}

// New creates a metadata service controller
func New(repo metadataRepositoryInterface) *Controller {
	return &Controller{repo: repo}
}

// Get returns movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
