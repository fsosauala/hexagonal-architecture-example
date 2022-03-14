package adapters

import (
	"context"

	"github.com/fsosauala/lambda-example/internal/core/domain"
)

type CountryRepository struct {
}

func NewCountryRepository() CountryRepository {
	return CountryRepository{}
}

func (cr CountryRepository) CreateCountry(ctx context.Context, country domain.Country) error {
	return nil
}
