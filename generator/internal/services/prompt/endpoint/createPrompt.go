package endpoint

import (
	"context"
	"net/http"

	"generator/internal/services/prompt/model"

	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Создание промпта по фильтрам
// @Tags prompt
// @Param Body body model.CreatePromptReq false "model.CreatePromptReq"
// @Produce json
// @Success 200 {object} model.CreatePromptRes
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /prompt [post]
func (e *endpoint) createPrompt(ctx context.Context, r *http.Request) (any, error) {

	var req model.CreatePromptReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.CreatePrompt(ctx, req)
}
