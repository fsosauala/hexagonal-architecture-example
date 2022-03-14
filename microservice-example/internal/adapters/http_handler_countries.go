package adapters

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fsosauala/microservice-example/internal/core/domain"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetCountries(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.countriesService.GetCountries(ctx)
	if err != nil {
		errorToReturn := domain.ErrUnknownError
		var ce domain.CustomErr
		if errors.As(err, &ce) {
			errorToReturn = ce
		}
		return c.JSON(errorToReturn.HTTPCode, errorToReturn)
	}

	return c.JSON(http.StatusOK, response)
}

func (h Handler) CreateCountry(c echo.Context) error {
	var request domain.CountryRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(domain.ErrCannotParseBody.HTTPCode, domain.ErrCannotParseBody)
	}

	ctx := c.Request().Context()
	response, err := h.countriesService.CreateCountry(ctx, request)
	if err != nil {
		errorToReturn := domain.ErrUnknownError
		var ce domain.CustomErr
		if errors.As(err, &ce) {
			errorToReturn = ce
		}
		return c.JSON(errorToReturn.HTTPCode, errorToReturn)
	}

	return c.JSON(http.StatusCreated, response)
}
