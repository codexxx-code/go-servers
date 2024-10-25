package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"pkg/http/chain"
	forecastService "server/internal/services/forecast/service"

	"server/internal/services/forecast/model"
)

type endpoint struct {
	service ForecastService
}

var _ ForecastService = new(forecastService.ForecastService)

type ForecastService interface {
	GetForecasts(context.Context, model.GetForecastsReq) ([]model.Forecast, error)
}

// MountForecastEndpoints mounts forecast endpoints to the router
func MountForecastEndpoints(mux *chi.Mux, service ForecastService) {
	mux.Mount("/forecast", newForecastEndpoint(service))
}

func newForecastEndpoint(service ForecastService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", chain.NewChain(e.getForecasts)) // GET /forecast

	return r
}
