package ports

import (
	"context"
	"github.com/fsosauala/lambda-example/internal/core/domain"
)

//go:generate mockgen -destination=../mocks/repositories_ports_mock.go -package=mocks -source=repositories_ports.go

type CountriesRepositoryPort interface {
	CreateCountry(ctx context.Context, country domain.Country) error
}
