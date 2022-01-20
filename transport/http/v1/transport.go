package v1

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/go-kit/kit/examples/addsvc/pkg/addservice"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"

	"net/http"
)

type gRPCServer struct {
	add gt.Handler
	*sync.Mutex
}

func NewHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	if zipkinTracer != nil {
		options = append(options, zipkin.HTTPServerTrace(zipkinTracer))
	}

	m := http.NewServeMux()

	m.Handle("/create-beer", httptransport.NewServer(
		endpoints.CreateBeerEndpoint,
		decodeHTTPCreateBeerRequest,
		encodeHTTPGenericResponse,
		options..., //append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "", logger)))...,
	))
	m.Handle("/update-beer", httptransport.NewServer(
		endpoints.UpdateBeerEndpoint,
		decodeHTTPUpdateBeerRequest,
		encodeHTTPGenericResponse,
		options..., //append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "", logger)))...,
	))
	m.Handle("/get-beer", httptransport.NewServer(
		endpoints.GetOneBeerEndpoint,
		decodeHTTPGetBeerRequest,
		encodeHTTPGenericResponse,
		options..., //append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "", logger)))...,
	))

	m.Handle("/list-beer", httptransport.NewServer(
		endpoints.ListBeersEndPoint,
		decodeHTTPListBeerRequest,
		encodeHTTPGenericResponse,
		options..., //append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "", logger)))...,
	))
	m.Handle("/delete-beer", httptransport.NewServer(
		endpoints.DeleteBeerEndpoint,
		decodeHTTPDeleteSecretRequest,
		encodeHTTPGenericResponse,
		options..., //append(options, httptransport.ServerBefore(
	// .HTTPToContext(otTracer, "", logger)))...,
	))
	return m
}
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case addservice.ErrTwoZeroes, addservice.ErrMaxSizeExceeded, addservice.ErrIntOverflow:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}

func decodeHTTPCreateBeerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req beerv1alpha1.CreateBeerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPListBeerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req beerv1alpha1.ListBeersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPDeleteSecretRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req beerv1alpha1.DeleteBeerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPUpdateBeerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req beerv1alpha1.UpdateBeerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPGetBeerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req beerv1alpha1.GetOneBeerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(GenericResponse); ok && f.Error != nil {
		errorEncoder(ctx, f.Error, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
