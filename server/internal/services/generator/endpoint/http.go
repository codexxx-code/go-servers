package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"pkg/http/chain"
	generatorModel "server/internal/services/generator/model"
)

type endpoint struct {
	service GeneratorService
}

type GeneratorService interface {
	GenerateDailyZodiacForecast(ctx context.Context, req generatorModel.GenerateDailyZodiacForecastReq) error
}

// MountGeneratorEndpoints mounts generator endpoints to the router
func MountGeneratorEndpoints(mux *chi.Mux, service GeneratorService) {
	mux.Mount("/generator", newGeneratorEndpoint(service))
}

func newGeneratorEndpoint(service GeneratorService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/dailyZodiacForecast", chain.NewChain(e.generateDailyZodiacForecast)) // GET /generator/dailyZodiacForecast

	return r
}
