package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/prompt/model"
	"generator/internal/services/prompt/repository/promptDDL"

	"pkg/errors"
)

func (r *PromptRepository) UpdatePrompt(ctx context.Context, req model.UpdatePromptReq) error {

	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if req.Case != nil {
		updates[promptDDL.ColumnCase] = *req.Case
	}
	if req.Text != nil {
		updates[promptDDL.ColumnText] = *req.Text
	}
	if req.Language != nil {
		updates[promptDDL.ColumnLanguage] = *req.Language
	}

	// Проверяем, что есть поля для обновления
	if len(updates) == 0 {
		return errors.BadRequest.New("No req to update")
	}

	// Создаем транзакцию
	return r.db.Exec(ctx, sq.
		Update(promptDDL.Table).
		SetMap(updates).
		Where(sq.Eq{promptDDL.ColumnID: req.ID}),
	)

}
