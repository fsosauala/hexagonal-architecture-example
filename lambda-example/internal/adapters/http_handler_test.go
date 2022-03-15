package adapters

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fsosauala/lambda-example/internal/core/domain"
	"github.com/fsosauala/lambda-example/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestHandler_ProcessRequest(t *testing.T) {
	type fields struct {
		countriesService func(m *mocks.MockCountriesServicePort)
	}
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "fail: cannot parse body",
			args: args{
				req: events.APIGatewayProxyRequest{
					Body: ``,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: domain.ErrCannotParseBody.HTTPCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: domain.ErrCannotParseBody.String(),
			},
			wantErr: false,
		},
		{
			name: "fail: countriesService.CreateCountry returns an error",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().CreateCountry(gomock.Any(), domain.CountryRequest{
						Name: "Argentina",
					}).Return(domain.CountryResponse{}, domain.ErrUnknownError)
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					Body: `{"name":"Argentina"}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: domain.ErrUnknownError.HTTPCode,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: domain.ErrUnknownError.String(),
			},
			wantErr: false,
		},
		{
			name: "success: full right test",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().CreateCountry(gomock.Any(), domain.CountryRequest{
						Name: "Argentina",
					}).Return(
						domain.CountryResponse{
							ID: "12345",
						},
						nil,
					)
				},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					Body: `{"name":"Argentina"}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusCreated,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"id":"12345"}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCountriesService := mocks.NewMockCountriesServicePort(mockCtrl)
			if tt.fields.countriesService != nil {
				tt.fields.countriesService(mockCountriesService)
			}

			h := Handler{
				countriesService: mockCountriesService,
			}
			got, err := h.ProcessRequest(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
