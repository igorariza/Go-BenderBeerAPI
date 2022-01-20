package data

import (
	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"

	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	beerv1alpha1.UnimplementedBeerAPIServiceServer
}
