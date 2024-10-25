package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"
	"pkg/validator"
	"server/internal/services/prompt/model"
)

// @Summary Удаление промпта
// @Tags prompt
// @Param Query query model.DeletePromptReq false "model.DeletePromptReq"
// @Produce json
// @Success 200
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /prompt [delete]
func (e *endpoint) deletePrompt(ctx context.Context, r *http.Request) (any, error) {

	var req model.DeletePromptReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, e.service.DeletePrompt(ctx, req)
}
