package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	generatorModel "generator/internal/services/generator/model"
	generatorService "generator/internal/services/generator/service"

	"pkg/http/chain"
)

type endpoint struct {
	service GeneratorService
}

var _ GeneratorService = new(generatorService.GeneratorService)

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
