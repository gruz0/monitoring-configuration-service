package configuration

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetConfigurationsEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a configuration
// server.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetConfigurationsEndpoint: MakeGetConfigurationsEndpoint(s),
	}
}

// GetConfiguration implements Service. Primarily useful in a client.
func (e Endpoints) GetConfigurations(ctx context.Context) ([]Configuration, error) {
	request := getConfigurationsRequest{}
	response, err := e.GetConfigurationsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getConfigurationsResponse)
	return []Configuration{}, resp.Err
	// return resp.Configuration, resp.Err
}

func MakeGetConfigurationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(getConfigurationsRequest)
		p, e := s.GetConfigurations(ctx)
		return getConfigurationsResponse{Configurations: p, Err: e}, nil
	}
}

type getConfigurationsRequest struct {
	// ID string
}

type getConfigurationsResponse struct {
	Configurations []Configuration `json:"profiles,omitempty"`
	Err            error           `json:"err,omitempty"`
}

func (r getConfigurationsResponse) error() error { return r.Err }
