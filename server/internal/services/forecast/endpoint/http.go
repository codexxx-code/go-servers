package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"pkg/http/chain"

	"server/internal/services/forecast/model"
)

type endpoint struct {
	service forecastService
}

type forecastService interface {
	GetForecasts(context.Context, model.GetForecastsReq) ([]model.Forecast, error)
}

// MountForecastEndpoints mounts forecast endpoints to the router
func MountForecastEndpoints(mux *chi.Mux, service forecastService) {
	mux.Mount("/forecast", newForecastEndpoint(service))
}

func newForecastEndpoint(service forecastService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", chain.NewChain(e.getForecasts)) // GET /forecast

	return r
}
