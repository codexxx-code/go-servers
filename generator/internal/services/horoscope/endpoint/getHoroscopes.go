package endpoint

import (
	"context"
	"net/http"

	"pkg/http/decoder"
	"pkg/validator"

	"generator/internal/services/horoscope/endpoint/model"
)

// @Summary Получение гороскопов по знакам зодиака
// @Tags horoscope
// @Param Query query model.GetHoroscopesReq false "model.GetHoroscopesReq"
// @Produce json
// @Success 200 {object} []model.Horoscope
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /horoscope [get]
func (e *endpoint) getHoroscopes(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetHoroscopesReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	businessModel, err := req.ConvertToBusinessModel()
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetHoroscopes(ctx, businessModel)
}
