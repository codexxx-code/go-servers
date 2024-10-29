package service

import (
	"context"

	"generator/internal/services/prompt/model"
)

func (s *PromptService) GetPrompts(ctx context.Context, req model.GetPromptsReq) (prompts []model.Prompt, err error) {
	return s.promptRepository.GetPrompts(ctx, req)
}
