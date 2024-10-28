package service

import (
	"context"

	"generator/internal/services/prompt/model"
)

func (s *PromptService) DeletePrompt(ctx context.Context, req model.DeletePromptReq) error {
	return s.promptRepository.DeletePrompt(ctx, req)
}
