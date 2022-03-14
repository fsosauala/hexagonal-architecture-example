package container

import (
	"context"

	"github.com/fsosauala/lambda-example/internal/core/services"

	"github.com/fsosauala/lambda-example/internal/adapters"

	"github.com/aws/aws-lambda-go/events"
)

type lambdaFunc func(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (response events.APIGatewayProxyResponse, err error)

type LambdaHandler struct {
	httpHandlerFunc lambdaFunc
}

func Initialize() LambdaHandler {
	countriesRepository := adapters.NewCountryRepository()
	countriesService := services.NewCountryService(countriesRepository)
	httpHandler := adapters.NewHTTPHandler(countriesService)
	return LambdaHandler{
		httpHandlerFunc: httpHandler.ProcessRequest,
	}
}

func (lambda *LambdaHandler) LambdaHandler(
	ctx context.Context, req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	return lambda.httpHandlerFunc(ctx, req)
}
