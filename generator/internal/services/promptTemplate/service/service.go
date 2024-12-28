package service

import (
	"context"

	promptModel "generator/internal/services/promptTemplate/model"
	promptRepository "generator/internal/services/promptTemplate/repository"
)

type PromptTemplateService struct {
	promptTemplateRepository PromptTemplateRepository
}

var _ PromptTemplateRepository = new(promptRepository.PromptTemplateRepository)

type PromptTemplateRepository interface {
	GetPromptTemplates(ctx context.Context, req promptModel.GetPromptTemplatesReq) ([]promptModel.PromptTemplate, error)
	UpdatePromptTemplate(ctx context.Context, req promptModel.UpdatePromptTemplateReq) error
}

func NewPromptTemplateService(
	promptRepository PromptTemplateRepository,
) *PromptTemplateService {
	return &PromptTemplateService{
		promptTemplateRepository: promptRepository,
	}
}
