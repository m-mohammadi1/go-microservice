package movie

import (
	"context"
	"errors"

	metadatamodel "movieexample.com/metadata/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg"
)

// ErrNotFound not found error
var ErrNotFound = errors.New("not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a movie service controller.
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{
		ratingGateway,
		metadataGateway,
	}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)

	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		// proceed
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
