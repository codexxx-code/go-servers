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

	// Categories
	GetCategories(context.Context, model.GetCategoriesReq) ([]model.Category, error)
	GetBreadcrumbs(context.Context, model.GetBreadcrumbsReq) (model.Category, error)

	// Items
	GetItems(context.Context, model.GetItemsReq) ([]model.Item, error)
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

	// Categories
	r.Method(http.MethodGet, "/categories", chain.NewChain(e.getCategories))            // GET /ebay/categories
	r.Method(http.MethodGet, "/category/breadcrumbs", chain.NewChain(e.getBreadcrumbs)) // GET /ebay/category/breadcrumbs

	// Items
	r.Method(http.MethodGet, "/items", chain.NewChain(e.getItems)) // GET /ebay/items

	return r
}
