package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/promptTemplate/model"
	"generator/internal/services/promptTemplate/repository/promptTemplateDDL"

	"pkg/ddlHelper"
)

func (r *PromptTemplateRepository) GetPromptTemplates(ctx context.Context, req model.GetPromptTemplatesReq) (prompts []model.PromptTemplate, err error) {

	filters := make(sq.Eq)

	if len(req.Cases) != 0 {
		filters[promptTemplateDDL.ColumnCase] = req.Cases
	}

	// Выполняем запрос
	return prompts, r.db.Select(ctx, &prompts, sq.
		Select(ddlHelper.SelectAll).
		From(promptTemplateDDL.Table).
		Where(filters),
	)

}
