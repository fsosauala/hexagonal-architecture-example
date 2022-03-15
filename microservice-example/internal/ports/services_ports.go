package ports

import (
	"context"

	"github.com/fsosauala/microservice-example/internal/core/domain"
)

//go:generate mockgen -destination=../mocks/services_ports_mock.go -package=mocks -source=services_ports.go

type CountriesServicePort interface {
	CreateCountry(ctx context.Context, request domain.CountryRequest) (domain.CountryResponse, error)
	GetCountries(ctx context.Context) (domain.GetCountriesResponse, error)
}
