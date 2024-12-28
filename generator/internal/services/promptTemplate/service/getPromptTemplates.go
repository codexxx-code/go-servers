package service

import (
	"context"

	"generator/internal/services/promptTemplate/model"
)

func (s *PromptTemplateService) GetPromptTemplates(ctx context.Context, req model.GetPromptTemplatesReq) (prompts []model.PromptTemplate, err error) {
	return s.promptTemplateRepository.GetPromptTemplates(ctx, req)
}
