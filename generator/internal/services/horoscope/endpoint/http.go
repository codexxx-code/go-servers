package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	horoscopeService "generator/internal/services/horoscope/service"

	"pkg/http/chain"

	"generator/internal/services/horoscope/model"
)

type endpoint struct {
	service HoroscopeService
}

var _ HoroscopeService = new(horoscopeService.HoroscopeService)

type HoroscopeService interface {
	GetHoroscopes(context.Context, model.GetHoroscopesReq) ([]model.Horoscope, error)
	GetHoroscopeWithGeneration(context.Context, model.GetHoroscopeWithGenerationReq) (model.Horoscope, error)
}

// MountHoroscopeEndpoints mounts horoscope endpoints to the router
func MountHoroscopeEndpoints(mux *chi.Mux, service HoroscopeService) {
	mux.Mount("/horoscope", newHoroscopeEndpoint(service))
}

func newHoroscopeEndpoint(service HoroscopeService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", chain.NewChain(e.getHoroscopes))                            // GET /horoscope
	r.Method(http.MethodGet, "/withGeneration", chain.NewChain(e.getHoroscopeWithGeneration)) // GET /horoscope/withGeneration

	return r
}
