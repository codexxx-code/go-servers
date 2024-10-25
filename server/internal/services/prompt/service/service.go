package service

import (
	"context"

	promptModel "server/internal/services/prompt/model"
)

type PromptService struct {
	promptRepository PromptRepository
}

type PromptRepository interface {
	GetPrompts(ctx context.Context, req promptModel.GetPromptsReq) ([]promptModel.Prompt, error)
}

func NewPromptService(
	promptRepository PromptRepository,
) *PromptService {
	return &PromptService{
		promptRepository: promptRepository,
	}
}
