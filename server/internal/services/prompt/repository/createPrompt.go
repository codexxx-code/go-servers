package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/prompt/model"
	"server/internal/services/prompt/repository/promptDDL"
)

func (r *PromptRepository) CreatePrompt(ctx context.Context, req model.CreatePromptReq) (res model.CreatePromptRes, err error) {
	id, err := r.db.ExecWithLastInsertID(ctx, sq.Insert(promptDDL.Table).
		SetMap(map[string]any{
			promptDDL.ColumnCase:     req.Case,
			promptDDL.ColumnLanguage: req.Language,
			promptDDL.ColumnText:     req.Text,
		}),
	)
	if err != nil {
		return res, err
	}

	return model.CreatePromptRes{ID: id}, nil
}
