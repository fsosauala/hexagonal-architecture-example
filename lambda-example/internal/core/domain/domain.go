package domain

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
