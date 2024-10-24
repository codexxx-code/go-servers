package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"
	"pkg/validator"
	"server/internal/services/generator/model"
)

// @Summary Генерация прогноза на один день для одного знака зодиака
// @Tags generator
// @Param Query query model.GenerateDailyZodiacForecastReq false "model.GenerateDailyZodiacForecastReq"
// @Produce json
// @Success 200
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /generator/dailyZodiacForecast [get]
func (e *endpoint) generateDailyZodiacForecast(ctx context.Context, r *http.Request) (any, error) {

	var req model.GenerateDailyZodiacForecastReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, e.service.GenerateDailyZodiacForecast(ctx, req)
}
