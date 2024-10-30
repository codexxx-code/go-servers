package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"partners/internal/services/ebay/model"
	ebaySerivce "partners/internal/services/ebay/service"
	"pkg/http/chain"
)

type endpoint struct {
	service EbayService
}

var _ EbayService = new(ebaySerivce.EbayService)

type EbayService interface {
	GetCategories(ctx context.Context, req model.GetCategoriesReq) ([]model.Category, error)
}

// MountEbayEndpoints mounts ebay endpoints to the router
func MountEbayEndpoints(mux *chi.Mux, service EbayService) {
	mux.Mount("/ebay", newEbayEndpoint(service))
}

func newEbayEndpoint(service EbayService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()

	r.Method(http.MethodGet, "/category", chain.NewChain(e.getCategories))

	return r
}
