package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fsosauala/lambda-example/internal/core/domain"
	"github.com/fsosauala/lambda-example/internal/ports"
)

type Handler struct {
	countriesService ports.CountriesServicePort
}

func NewHTTPHandler(cs ports.CountriesServicePort) Handler {
	return Handler{
		countriesService: cs,
	}
}

func (h Handler) ProcessRequest(
	ctx context.Context, req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	var request domain.CountryRequest
	if err := json.NewDecoder(strings.NewReader(req.Body)).Decode(&request); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: domain.ErrCannotParseBody.HTTPCode,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: domain.ErrCannotParseBody.String(),
		}, nil
	}

	response, err := h.countriesService.CreateCountry(ctx, request)
	if err != nil {
		toReturn := events.APIGatewayProxyResponse{
			Body:       domain.ErrUnknownError.String(),
			StatusCode: domain.ErrUnknownError.HTTPCode,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		var ce domain.CustomErr
		if errors.As(err, &ce) {
			toReturn.StatusCode = ce.HTTPCode
			toReturn.Body = ce.String()
		}
		return toReturn, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: response.String(),
	}, nil
}
