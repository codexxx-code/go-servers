package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/prompt/model"
	"generator/internal/services/prompt/repository/promptDDL"

	"pkg/ddlHelper"
)

func (r *PromptRepository) GetPrompts(ctx context.Context, req model.GetPromptsReq) (prompts []model.Prompt, err error) {

	filters := make(sq.Eq)

	if len(req.Cases) != 0 {
		filters[promptDDL.ColumnCase] = req.Cases
	}

	if len(req.Languages) != 0 {
		filters[promptDDL.ColumnLanguage] = req.Languages
	}

	// Выполняем запрос
	return prompts, r.db.Select(ctx, &prompts, sq.
		Select(ddlHelper.SelectAll).
		From(promptDDL.Table).
		Where(filters),
	)

}
