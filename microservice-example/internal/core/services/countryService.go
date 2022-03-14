package services

import (
	"context"
	"errors"

	"github.com/fsosauala/microservice-example/internal/core/domain"
	"github.com/fsosauala/microservice-example/internal/ports"
	"github.com/google/uuid"
)

type CountryService struct {
	countriesRepository ports.CountriesRepositoryPort
}

func NewCountryService(cr ports.CountriesRepositoryPort) CountryService {
	return CountryService{
		countriesRepository: cr,
	}
}

func (cs CountryService) CreateCountry(ctx context.Context, request domain.CountryRequest) (domain.CountryResponse, error) {
	if request.Name == "" {
		return domain.CountryResponse{}, domain.ErrEmptyName
	}

	country := domain.Country{
		ID:   uuid.NewString(),
		Name: request.Name,
	}

	if err := cs.countriesRepository.CreateCountry(ctx, country); err != nil {
		if errors.Is(err, domain.AlreadyExistsError) {
			return domain.CountryResponse{}, domain.ErrBadRequest
		}

		return domain.CountryResponse{}, domain.ErrUnknownError
	}

	return domain.CountryResponse{
		ID: country.ID,
	}, nil
}
