package endpoint

import (
	"context"
	"net/http"

	"generator/internal/services/horoscope/model"
	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Получение гороскопа по знакам зодиака, если его нет - генерация
// @Tags horoscope
// @Param Query query model.GetHoroscopeWithGenerationReq false "model.GetHoroscopeWithGenerationReq"
// @Produce json
// @Success 200 {object} model.Horoscope
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /horoscope/withGeneration [get]
func (e *endpoint) getHoroscopeWithGeneration(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetHoroscopeWithGenerationReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetHoroscopeWithGeneration(ctx, req)
}
