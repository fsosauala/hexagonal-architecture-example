package adapters

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fsosauala/microservice-example/internal/core/domain"
	"github.com/fsosauala/microservice-example/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateCountry(t *testing.T) {
	type fields struct {
		countriesService func(m *mocks.MockCountriesServicePort)
	}
	type args struct {
		body string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		expectedCode int
		expectedBody string
	}{
		{
			name: "fail: invalid body",
			args: args{
				body: "",
			},
			wantErr:      false,
			expectedCode: domain.ErrCannotParseBody.HTTPCode,
			expectedBody: `{"message":"Error parsing body","errorCode":2}`,
		},
		{
			name: "fail: countriesService.CreateCountry returns an error",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().CreateCountry(gomock.Any(), domain.CountryRequest{
						Name: "Argentina",
					}).Return(domain.CountryResponse{}, domain.ErrBadRequest)
				},
			},
			args: args{
				body: `{"name":"Argentina"}`,
			},
			expectedCode: domain.ErrBadRequest.HTTPCode,
			expectedBody: `{"message":"The country already exists","errorCode":4}`,
		},
		{
			name: "success: full right test",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().CreateCountry(gomock.Any(), domain.CountryRequest{
						Name: "Argentina",
					}).Return(domain.CountryResponse{
						ID: "123456",
					}, nil)
				},
			},
			args: args{
				body: `{"name":"Argentina"}`,
			},
			wantErr:      false,
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":"123456"}`,
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

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.args.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := h.CreateCountry(c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCountry() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Equal(t, tt.expectedBody, strings.ReplaceAll(rec.Body.String(), "\n", ""))
		})
	}
}

func TestHandler_GetCountries(t *testing.T) {
	type fields struct {
		countriesService func(m *mocks.MockCountriesServicePort)
	}

	tests := []struct {
		name         string
		fields       fields
		wantErr      bool
		expectedCode int
		expectedBody string
	}{
		{
			name: "fail: countriesService.GetCountries returns an error",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().GetCountries(gomock.Any()).
						Return(domain.GetCountriesResponse{}, domain.ErrBadRequest)
				},
			},
			wantErr:      false,
			expectedCode: domain.ErrBadRequest.HTTPCode,
			expectedBody: `{"message":"The country already exists","errorCode":4}`,
		},
		{
			name: "success: full right test",
			fields: fields{
				countriesService: func(m *mocks.MockCountriesServicePort) {
					m.EXPECT().GetCountries(gomock.Any()).
						Return(domain.GetCountriesResponse{
							Countries: []domain.Country{
								{
									ID:   "123456",
									Name: "Argentina",
								},
							},
						}, nil)
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
			expectedBody: `{"countries":[{"id":"123456","name":"Argentina"}]}`,
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

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := h.GetCountries(c); (err != nil) != tt.wantErr {
				t.Errorf("GetCountries() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Equal(t, tt.expectedBody, strings.ReplaceAll(rec.Body.String(), "\n", ""))
		})
	}
}
