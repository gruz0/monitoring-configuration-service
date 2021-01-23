package configuration

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type listConfigurationsRequest struct{}

type listConfigurationsResponse struct {
	Configuration []Configuration `json:"configuration"`
}

func makeListConfigurationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listConfigurationsRequest)
		return listConfigurationsResponse{Configuration: s.Configurations()}, nil
	}
}
