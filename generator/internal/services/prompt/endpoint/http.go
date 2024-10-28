package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	promptModel "generator/internal/services/prompt/model"
	promptService "generator/internal/services/prompt/service"

	"pkg/http/chain"
)

type endpoint struct {
	service PromptService
}

var _ PromptService = new(promptService.PromptService)

type PromptService interface {
	GetPrompts(ctx context.Context, req promptModel.GetPromptsReq) ([]promptModel.Prompt, error)
	UpdatePrompt(ctx context.Context, req promptModel.UpdatePromptReq) error
	DeletePrompt(ctx context.Context, req promptModel.DeletePromptReq) error
	CreatePrompt(ctx context.Context, req promptModel.CreatePromptReq) (promptModel.CreatePromptRes, error)
}

// MountPromptEndpoints mounts prompt endpoints to the router
func MountPromptEndpoints(mux *chi.Mux, service PromptService) {
	mux.Mount("/prompt", newPromptEndpoint(service))
}

func newPromptEndpoint(service PromptService) http.Handler {

	e := &endpoint{
		service: service,
	}

	r := chi.NewRouter()
	r.Method(http.MethodGet, "/", chain.NewChain(e.getPrompts))      // GET /prompt
	r.Method(http.MethodPatch, "/", chain.NewChain(e.updatePrompt))  // PATCH /prompt
	r.Method(http.MethodDelete, "/", chain.NewChain(e.deletePrompt)) // DELETE /prompt
	r.Method(http.MethodPost, "/", chain.NewChain(e.createPrompt))   // POST /prompt

	return r
}
