package domain

import "encoding/json"

type (
	CountryRequest struct {
		Name string `json:"name"`
	}
	CountryResponse struct {
		ID string `json:"id"`
	}
	Country struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func (r CountryResponse) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(data)
}
