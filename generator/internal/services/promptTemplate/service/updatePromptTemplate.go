package service

import (
	"context"
	"text/template"

	"generator/internal/services/promptTemplate/model"
	"pkg/errors"
)

func (s *PromptTemplateService) UpdatePromptTemplate(ctx context.Context, req model.UpdatePromptTemplateReq) error {

	// Валидируем темплейт
	if req.Template != nil {
		if _, err := template.New("index").Parse(*req.Template); err != nil {
			return errors.BadRequest.Wrap(err)
		}
	}

	return s.promptTemplateRepository.UpdatePromptTemplate(ctx, req)
}
