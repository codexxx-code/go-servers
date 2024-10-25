package service

import (
	"context"

	"server/internal/services/prompt/model"
)

func (s *PromptService) DeletePrompt(ctx context.Context, req model.DeletePromptReq) error {
	return s.promptRepository.DeletePrompt(ctx, req)
}
