package ports

import (
	"context"
	"github.com/fsosauala/lambda-example/internal/core/domain"
)

type CountriesRepositoryPort interface {
	CreateCountry(ctx context.Context, country domain.Country) error
}
