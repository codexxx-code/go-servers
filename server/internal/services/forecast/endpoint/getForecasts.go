package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"

	"server/internal/services/forecast/model"
)

// @Summary Получение прогнозов по знакам зодиака
// @Tags forecast
// @Param Query query model.GetForecastsReq false "model.GetForecastsReq"
// @Produce json
// @Success 200 {object} []model.Forecast
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /forecast [get]
func (e *endpoint) getForecasts(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetForecastsReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetForecasts(ctx, req)
}
