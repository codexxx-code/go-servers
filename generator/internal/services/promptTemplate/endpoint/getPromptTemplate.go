package endpoint

import (
	"context"
	"net/http"

	"generator/internal/services/promptTemplate/model"

	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Получение темплейтов для формирования промптов по фильтрам
// @Tags prompt
// @Param Query query model.GetPromptTemplatesReq false "model.GetPromptTemplatesReq"
// @Produce json
// @Success 200 {object} []model.PromptTemplate
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /promptTemplate [get]
func (e *endpoint) getPromptTemplates(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetPromptTemplatesReq

	// Декодируем запрос
	if err := decoder.Decode(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetPromptTemplates(ctx, req)
}
