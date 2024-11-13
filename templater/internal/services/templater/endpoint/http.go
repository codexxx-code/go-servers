package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"templater/internal/services/templater/model"
	templaterSerivce "templater/internal/services/templater/service"

	"pkg/http/chain"
	"pkg/openrtb"
)

type endpoint struct {
	service TemplaterService
}

var _ TemplaterService = new(templaterSerivce.TemplaterService)

type TemplaterService interface {
	GetTemplate(context.Context, model.GetTemplateReq) (openrtb.BidResponse, error)
}

// MountTemplaterEndpoints mounts templater endpoints to the router
func MountTemplaterEndpoints(mux *chi.Mux, service TemplaterService) {
	mux.Mount("/rtb", newTemplaterEndpoint(service))
}

func newTemplaterEndpoint(service TemplaterService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()

	r.Method(http.MethodPost, "/{sspSlug}", chain.NewChain(e.getTemplate)) // POST /rtb/{sspSlug}

	return r
}
