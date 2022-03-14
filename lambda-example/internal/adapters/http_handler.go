package adapters

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
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
	return events.APIGatewayProxyResponse{}, nil
}
