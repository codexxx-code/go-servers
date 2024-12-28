package service

import (
	"context"

	"generator/internal/services/promptTemplate/model"
)

func (s *PromptTemplateService) UpdatePromptTemplate(ctx context.Context, req model.UpdatePromptTemplateReq) error {
	return s.promptTemplateRepository.UpdatePromptTemplate(ctx, req)
}
