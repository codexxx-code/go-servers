package endpoint

import (
	"net/http"

	"github.com/go-chi/chi"
)

type endpoint struct {
	service zodiacService
}

type zodiacService interface {
}

// MountZodiacEndpoints mounts zodiac endpoints to the router
func MountZodiacEndpoints(mux *chi.Mux, service zodiacService) {
	mux.Mount("/zodiac", newZodiacEndpoint(service))
}

func newZodiacEndpoint(service zodiacService) http.Handler {

	e := &endpoint{
		service: service,
	}

	_ = e

	r := chi.NewRouter()

	return r
}
