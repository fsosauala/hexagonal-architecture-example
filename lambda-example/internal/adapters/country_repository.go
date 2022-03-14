package adapters

import (
	"context"

	"github.com/fsosauala/lambda-example/internal/core/domain"
)

type CountryRepository struct {
	db map[string]domain.Country
}

func NewCountryRepository() CountryRepository {
	db := make(map[string]domain.Country)
	return CountryRepository{
		db: db,
	}
}

func (cr CountryRepository) CreateCountry(ctx context.Context, country domain.Country) error {
	if _, exists := cr.db[country.Name]; exists {
		return domain.AlreadyExistsError
	}
	cr.db[country.Name] = country
	return nil
}
