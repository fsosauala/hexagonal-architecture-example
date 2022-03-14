package services

import (
	"context"

	"github.com/fsosauala/lambda-example/internal/core/domain"
)

type CountryService struct {
}

func NewCountryService() CountryService {
	return CountryService{}
}

func (cs CountryService) CreateCountry(ctx context.Context, request domain.CountryRequest) (domain.CountryResponse, error) {
	return domain.CountryResponse{}, nil
}
