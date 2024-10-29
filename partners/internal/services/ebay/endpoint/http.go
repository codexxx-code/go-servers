package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	ebaySerivce "partners/internal/services/ebay/service"
)

type endpoint struct {
	service EbayService
}

var _ EbayService = new(ebaySerivce.EbayService)

type EbayService interface{}

// MountEbayEndpoints mounts ebay endpoints to the router
func MountEbayEndpoints(mux *chi.Mux, service EbayService) {
	mux.Mount("/ebay", newEbayEndpoint(service))
}

func newEbayEndpoint(service EbayService) http.Handler {

	_ = &endpoint{
		service: service,
	}

	r := chi.NewRouter()

	return r
}
