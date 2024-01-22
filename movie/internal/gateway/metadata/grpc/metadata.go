package grpc

import (
	"context"

	"movieexample.com/gen"
	"movieexample.com/internal/grpcutil"
	"movieexample.com/metadata/pkg/model"
	"movieexample.com/pkg/discovery"
)

// Gateway defines a movie metadata grpc gateway.
type Gateway struct {
	registry discovery.Registry
}

// New creates a movie metadata grpc gateway.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})
	if err != nil {
		return nil, err
	}

	return model.MetadataFromProto(resp.Metadata), nil
}
