package adapters

import (
	"net/http"

	"github.com/fsosauala/microservice-example/internal/ports"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	countriesService ports.CountriesServicePort
}

func NewHTTPHandler(cs ports.CountriesServicePort) *echo.Echo {
	h := Handler{
		countriesService: cs,
	}

	// Echo instance
	e := echo.New()
	h.InitRoutes(e)

	// Start server
	return e
}

func (h Handler) InitRoutes(e *echo.Echo) {
	// Routes
	e.GET("/", h.Hello)
	h.InitCountriesRoutes(e)
}

func (h Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Ualá!")
}

func (h Handler) InitCountriesRoutes(c *echo.Echo) {
	countriesGroup := c.Group("/countries")
	countriesGroup.GET("", h.GetCountries)
	countriesGroup.POST("", h.CreateCountry)
}
