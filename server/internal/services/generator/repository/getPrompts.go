package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"server/internal/services/generator/model"
)

func (r *GeneratorRepository) GetPrompts(ctx context.Context, req model.GetPromptsReq) (prompts []model.Prompt, err error) {

	filters := make(sq.Eq)

	filters[`"case"`] = req.Case

	// Выполняем запрос
	return prompts, r.db.Select(ctx, &prompts, sq.
		Select(ddlHelper.SelectAll).
		From("zodiac.prompts").
		Where(filters),
	)

}
