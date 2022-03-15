package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/fsosauala/microservice-example/internal/core/domain"
	"github.com/fsosauala/microservice-example/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestCountryService_CreateCountry(t *testing.T) {
	type fields struct {
		countriesRepository func(m *mocks.MockCountriesRepositoryPort)
	}
	type args struct {
		request domain.CountryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.CountryResponse
		wantErr bool
	}{
		{
			name: "fail: empty name",
			args: args{
				request: domain.CountryRequest{
					Name: "",
				},
			},
			want:    domain.CountryResponse{},
			wantErr: true,
		},
		{
			name: "fail: countriesRepository.CreateCountry returns domain.AlreadyExistsError",
			fields: fields{
				countriesRepository: func(m *mocks.MockCountriesRepositoryPort) {
					m.EXPECT().CreateCountry(gomock.Any(), gomock.AssignableToTypeOf(domain.Country{})).
						Return(domain.AlreadyExistsError)
				},
			},
			args: args{
				request: domain.CountryRequest{
					Name: "Colombia",
				},
			},
			want:    domain.CountryResponse{},
			wantErr: true,
		},
		{
			name: "fail: countriesRepository.CreateCountry returns an unknown error",
			fields: fields{
				countriesRepository: func(m *mocks.MockCountriesRepositoryPort) {
					m.EXPECT().CreateCountry(gomock.Any(), gomock.AssignableToTypeOf(domain.Country{})).
						Return(fmt.Errorf("some error"))
				},
			},
			args: args{
				request: domain.CountryRequest{
					Name: "Colombia",
				},
			},
			want:    domain.CountryResponse{},
			wantErr: true,
		},
		{
			name: "success: full right test",
			fields: fields{
				countriesRepository: func(m *mocks.MockCountriesRepositoryPort) {
					m.EXPECT().CreateCountry(gomock.Any(), gomock.AssignableToTypeOf(domain.Country{})).
						Return(nil)
				},
			},
			args: args{
				request: domain.CountryRequest{
					Name: "Colombia",
				},
			},
			want: domain.CountryResponse{
				ID: "pending",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCountryRepository := mocks.NewMockCountriesRepositoryPort(mockCtrl)
			if tt.fields.countriesRepository != nil {
				tt.fields.countriesRepository(mockCountryRepository)
			}

			cs := CountryService{
				countriesRepository: mockCountryRepository,
			}
			got, err := cs.CreateCountry(context.Background(), tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				tt.want.ID = got.ID
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCountry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountryService_GetCountries(t *testing.T) {
	type fields struct {
		countriesRepository func(m *mocks.MockCountriesRepositoryPort)
	}

	tests := []struct {
		name    string
		fields  fields
		want    domain.GetCountriesResponse
		wantErr bool
	}{
		{
			name: "fail: countriesRepository.GetCountries returns an error",
			fields: fields{
				countriesRepository: func(m *mocks.MockCountriesRepositoryPort) {
					m.EXPECT().GetCountries(gomock.Any()).
						Return([]domain.Country{}, domain.ErrUnknownError)
				},
			},
			want:    domain.GetCountriesResponse{},
			wantErr: true,
		},
		{
			name: "success: full right test",
			fields: fields{
				countriesRepository: func(m *mocks.MockCountriesRepositoryPort) {
					m.EXPECT().GetCountries(gomock.Any()).
						Return([]domain.Country{
							{
								ID:   "12456",
								Name: "Argentina",
							},
						}, nil)
				},
			},
			want: domain.GetCountriesResponse{
				Countries: []domain.Country{
					{
						ID:   "12456",
						Name: "Argentina",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockCountryRepository := mocks.NewMockCountriesRepositoryPort(mockCtrl)
			if tt.fields.countriesRepository != nil {
				tt.fields.countriesRepository(mockCountryRepository)
			}

			cs := CountryService{
				countriesRepository: mockCountryRepository,
			}
			got, err := cs.GetCountries(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCountries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCountries() got = %v, want %v", got, tt.want)
			}
		})
	}
}
