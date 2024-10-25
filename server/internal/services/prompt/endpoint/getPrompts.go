package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"
	"pkg/validator"
	"server/internal/services/prompt/model"
)

// @Summary Получение промптов по фильтрам
// @Tags prompt
// @Param Query query model.GetPromptsReq false "model.GetPromptsReq"
// @Produce json
// @Success 200 {object} []model.Prompt
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /prompt [get]
func (e *endpoint) getPrompts(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetPromptsReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetPrompts(ctx, req)
}
