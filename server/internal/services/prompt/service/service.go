package service

import (
	"context"

	promptModel "server/internal/services/prompt/model"
	promptRepository "server/internal/services/prompt/repository"
)

type PromptService struct {
	promptRepository PromptRepository
}

var _ PromptRepository = new(promptRepository.PromptRepository)

type PromptRepository interface {
	CreatePrompt(ctx context.Context, req promptModel.CreatePromptReq) (promptModel.CreatePromptRes, error)
	GetPrompts(ctx context.Context, req promptModel.GetPromptsReq) ([]promptModel.Prompt, error)
	UpdatePrompt(ctx context.Context, req promptModel.UpdatePromptReq) error
	DeletePrompt(ctx context.Context, req promptModel.DeletePromptReq) error
}

func NewPromptService(
	promptRepository PromptRepository,
) *PromptService {
	return &PromptService{
		promptRepository: promptRepository,
	}
}
