package endpoint

import (
	"context"
	"net/http"

	"generator/internal/services/promptTemplate/model"

	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Обновление темплейта для формирования промпта (patch)
// @Tags prompt
// @Param Body body model.UpdatePromptTemplateReq false "model.UpdatePromptTemplateReq"
// @Produce json
// @Success 200
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /promptTemplate [patch]
func (e *endpoint) updatePromptTemplate(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdatePromptTemplateReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, e.service.UpdatePromptTemplate(ctx, req)
}
