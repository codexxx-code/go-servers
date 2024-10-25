package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"pkg/http/chain"
	promptModel "server/internal/services/prompt/model"
)

type endpoint struct {
	service PromptService
}

type PromptService interface {
	GetPrompts(ctx context.Context, req promptModel.GetPromptsReq) ([]promptModel.Prompt, error)
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
	r.Method(http.MethodGet, "/", chain.NewChain(e.getPrompts)) // GET /prompt

	return r
}
