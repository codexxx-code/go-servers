package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"
	"pkg/validator"
	"server/internal/services/prompt/model"
)

// @Summary Обновление промпта (patch)
// @Tags prompt
// @Param Body body model.UpdatePromptReq false "model.UpdatePromptReq"
// @Produce json
// @Success 200
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /prompt [patch]
func (e *endpoint) updatePrompt(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdatePromptReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, e.service.UpdatePrompt(ctx, req)
}
