package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"server/internal/services/prompt/model"
)

func (r *PromptRepository) GetPrompts(ctx context.Context, req model.GetPromptsReq) (prompts []model.Prompt, err error) {

	filters := make(sq.Eq)

	if len(req.Cases) != 0 {
		filters[`"case"`] = req.Cases
	}

	// Выполняем запрос
	return prompts, r.db.Select(ctx, &prompts, sq.
		Select(ddlHelper.SelectAll).
		From("zodiac.prompts").
		Where(filters),
	)

}
