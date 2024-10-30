package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/prompt/model"
	"generator/internal/services/prompt/repository/promptDDL"
)

func (r *PromptRepository) DeletePrompt(ctx context.Context, req model.DeletePromptReq) error {
	return r.db.Exec(ctx, sq.
		Delete(promptDDL.Table).
		Where(sq.Eq{promptDDL.ColumnID: req.ID}),
	)
}
