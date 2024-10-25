package service

import (
	"context"

	"server/internal/services/prompt/model"
)

func (s *PromptService) GetPrompts(ctx context.Context, req model.GetPromptsReq) (prompts []model.Prompt, err error) {
	return s.promptRepository.GetPrompts(ctx, req)
}
