package v1

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	svcbeerv1alpha1 "github.com/igorariza/gogrpc-apibeer/cmd/beer"
	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
)

type Endpoints struct {
	CreateBeerEndpoint endpoint.Endpoint
	ListBeersEndPoint  endpoint.Endpoint
	DeleteBeerEndpoint endpoint.Endpoint
	UpdateBeerEndpoint endpoint.Endpoint
	GetOneBeerEndpoint endpoint.Endpoint
}
type GenericRequest struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error,omitempty"`
}
type GenericResponse struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error,omitempty"`
}

func MakeEndpoints(s svcbeerv1alpha1.BeerApiService) Endpoints {
	return Endpoints{
		CreateBeerEndpoint: MakeCreateBeerEndpointEndpoint(s),
		ListBeersEndPoint:  MakeListBeersEndPointEndPoint(s),
		DeleteBeerEndpoint: MakeDeleteBeerEndpointEndpoint(s),
		UpdateBeerEndpoint: MakeUpdateBeerEndpointEndpoint(s),
		GetOneBeerEndpoint: MakeGetOneBeerEndpointEndpoint(s),
	}
}

func MakeCreateBeerEndpointEndpoint(s svcbeerv1alpha1.BeerApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(beerv1alpha1.CreateBeerRequest)
		result, err := s.CreateBeer(ctx, &req)
		return GenericResponse{Data: result, Error: err}, nil
	}
}

func MakeListBeersEndPointEndPoint(s svcbeerv1alpha1.BeerApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(beerv1alpha1.ListBeersRequest)
		result, err := s.ListBeers(ctx, &req)
		return GenericResponse{Data: result, Error: err}, nil
	}
}
func MakeDeleteBeerEndpointEndpoint(s svcbeerv1alpha1.BeerApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(beerv1alpha1.DeleteBeerRequest)
		result, err := s.DeleteBeer(ctx, &req)
		return GenericResponse{Data: result, Error: err}, nil
	}
}
func MakeUpdateBeerEndpointEndpoint(s svcbeerv1alpha1.BeerApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(beerv1alpha1.UpdateBeerRequest)
		result, err := s.UpdateBeer(ctx, &req)
		return GenericResponse{Data: result, Error: err}, nil
	}
}
func MakeGetOneBeerEndpointEndpoint(s svcbeerv1alpha1.BeerApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(beerv1alpha1.GetOneBeerRequest)
		result, err := s.GetBeer(ctx, &req)
		return GenericResponse{Data: result, Error: err}, nil
	}
}
