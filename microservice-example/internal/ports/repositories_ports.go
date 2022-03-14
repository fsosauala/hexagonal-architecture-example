package ports

import (
	"context"

	"github.com/fsosauala/microservice-example/internal/core/domain"
)

type CountriesRepositoryPort interface {
	CreateCountry(ctx context.Context, country domain.Country) error
}
