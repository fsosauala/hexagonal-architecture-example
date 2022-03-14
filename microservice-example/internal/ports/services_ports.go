package ports

import (
	"context"

	"github.com/fsosauala/microservice-example/internal/core/domain"
)

type CountriesServicePort interface {
	CreateCountry(ctx context.Context, request domain.CountryRequest) (domain.CountryResponse, error)
}
