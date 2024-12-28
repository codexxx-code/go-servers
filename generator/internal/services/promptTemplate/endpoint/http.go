package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	promptModel "generator/internal/services/promptTemplate/model"
	promptService "generator/internal/services/promptTemplate/service"

	"pkg/http/chain"
)

type endpoint struct {
	service PromptService
}

var _ PromptService = new(promptService.PromptTemplateService)

type PromptService interface {
	GetPromptTemplates(ctx context.Context, req promptModel.GetPromptTemplatesReq) ([]promptModel.PromptTemplate, error)
	UpdatePromptTemplate(ctx context.Context, req promptModel.UpdatePromptTemplateReq) error
}

// MountPromptTemplateEndpoints mounts prompt endpoints to the router
func MountPromptTemplateEndpoints(mux *chi.Mux, service PromptService) {
	mux.Mount("/promptTemplate", newPromptEndpoint(service))
}

func newPromptEndpoint(service PromptService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", chain.NewChain(e.getPromptTemplates))     // GET /promptTemplate
	r.Method(http.MethodPatch, "/", chain.NewChain(e.updatePromptTemplate)) // PATCH /promptTemplate

	return r
}
