package beer

import (
	"context"

	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
)

type Repository interface {
	Create(ctx context.Context, in *beerv1alpha1.CreateBeerRequest) (*beerv1alpha1.CreateBeerResponse, error)
	List(ctx context.Context, in *beerv1alpha1.ListBeersRequest) (*beerv1alpha1.ListBeersResponse, error)
	Update(ctx context.Context, in *beerv1alpha1.UpdateBeerRequest) (*beerv1alpha1.UpdateBeerResponse, error)
	Get(ctx context.Context, in *beerv1alpha1.GetOneBeerRequest) (*beerv1alpha1.GetOneBeerResponse, error)
	Delete(ctx context.Context, in *beerv1alpha1.DeleteBeerRequest) (*beerv1alpha1.DeleteBeerResponse, error)
}
