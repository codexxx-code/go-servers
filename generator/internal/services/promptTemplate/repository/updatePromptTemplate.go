package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/promptTemplate/model"
	"generator/internal/services/promptTemplate/repository/promptTemplateDDL"

	"pkg/errors"
)

func (r *PromptTemplateRepository) UpdatePromptTemplate(ctx context.Context, req model.UpdatePromptTemplateReq) error {

	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if req.Template != nil {
		updates[promptTemplateDDL.ColumnText] = *req.Template
	}

	// Проверяем, что есть поля для обновления
	if len(updates) == 0 {
		return errors.BadRequest.New("No req to update")
	}

	// Создаем транзакцию
	return r.db.Exec(ctx, sq.
		Update(promptTemplateDDL.Table).
		SetMap(updates).
		Where(sq.Eq{promptTemplateDDL.ColumnCase: req.Case}),
	)

}
