package svcbeerv1alpha1

import (
	"time"

	"context"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
	"github.com/igorariza/gogrpc-apibeer/internal/data"
	"github.com/igorariza/gogrpc-apibeer/pkg/beer"
)

type serverContext struct {
	// default timeout
	timeout time.Duration
	// some other useful objects, like config
	// or logger (to replace global logging)
	// (...)
}

type BeerApiService struct {
	beerv1alpha1.UnimplementedBeerAPIServiceServer
	Repository beer.Repository
	context    *serverContext
	logger     log.Logger
}

func (s *BeerApiService) CreateBeer(ctx context.Context, in *beerv1alpha1.CreateBeerRequest) (*beerv1alpha1.CreateBeerResponse, error) {
	s.logger.Log("method", "CreateBeer", "Creating Beer", in)
	response := &beerv1alpha1.CreateBeerResponse{
		Msg: "",
	}
	ber, err := s.Repository.Create(ctx, in)
	if err == nil {
		response.Msg = "Beer created"
		s.logger.Log("method", "CreateBeer", "status", ber)
		return response, nil
	} else {
		response.Msg = fmt.Sprint(err)
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, err
	}
}

func (s *BeerApiService) UpdateBeer(ctx context.Context, in *beerv1alpha1.UpdateBeerRequest) (*beerv1alpha1.UpdateBeerResponse, error) {
	s.logger.Log("method", "CreateBeer", "Creating Beer %v", in)
	response := &beerv1alpha1.UpdateBeerResponse{
		Msg: "",
	}
	_, err := s.Repository.Update(ctx, in)
	if err == nil {
		response.Msg = "Beer created"
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, nil
	} else {
		response.Msg = fmt.Sprint(err)
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, err
	}
}

func (s *BeerApiService) ListBeer(ctx context.Context, in *beerv1alpha1.ListBeersRequest) (*beerv1alpha1.ListBeersResponse, error) {
	s.logger.Log("method", "CreateBeer", "Creating Beer %v", in)
	response := &beerv1alpha1.ListBeersResponse{
		Msg: "",
	}
	_, err := s.Repository.List(ctx, in)
	if err == nil {
		response.Msg = "Beer created"
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, nil
	} else {
		response.Msg = fmt.Sprint(err)
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, err
	}
}
func (s *BeerApiService) GetBeer(ctx context.Context, in *beerv1alpha1.GetOneBeerRequest) (*beerv1alpha1.GetOneBeerResponse, error) {
	s.logger.Log("method", "CreateBeer", "Creating Beer %v", in)
	response := &beerv1alpha1.GetOneBeerResponse{
		Msg: "",
	}
	_, err := s.Repository.Get(ctx, in)
	if err == nil {
		response.Msg = "Beer created"
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, nil
	} else {
		response.Msg = fmt.Sprint(err)
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, err
	}
}

func (s *BeerApiService) DeleteBeer(ctx context.Context, in *beerv1alpha1.DeleteBeerRequest) (*beerv1alpha1.DeleteBeerResponse, error) {
	s.logger.Log("method", "CreateBeer", "Creating Beer %v", in)
	response := &beerv1alpha1.DeleteBeerResponse{
		Msg: "",
	}
	_, err := s.Repository.Delete(ctx, in)
	if err == nil {
		response.Msg = "Beer created"
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, nil
	} else {
		response.Msg = fmt.Sprint(err)
		s.logger.Log("method", "CreateBeer", "status", response.Msg)
		return response, err
	}
}

// NewService returns a Service with all the expected dependencies
func NewService(logger log.Logger) BeerApiService {
	fmt.Println("BeerAPIService instanced...")
	return BeerApiService{
		Repository: &data.BeerRepository{
			Data: data.New(),
		},
		context: &serverContext{
			timeout: time.Second,
		},
		logger: log.NewLogfmtLogger(os.Stderr),
	}
}
