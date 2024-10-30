package service

import (
	"context"

	"generator/internal/services/prompt/model"
)

func (s *PromptService) CreatePrompt(ctx context.Context, req model.CreatePromptReq) (model.CreatePromptRes, error) {
	return s.promptRepository.CreatePrompt(ctx, req)
}
