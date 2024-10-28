package service

import (
	"context"

	"generator/internal/services/prompt/model"
)

func (s *PromptService) UpdatePrompt(ctx context.Context, req model.UpdatePromptReq) error {
	return s.promptRepository.UpdatePrompt(ctx, req)
}
