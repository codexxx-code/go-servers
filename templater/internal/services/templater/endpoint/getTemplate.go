package endpoint

import (
	"context"
	"net/http"

	endpointModel "templater/internal/services/templater/endpoint/model"

	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Bidding
// @Tags rtb
// @Param sspSlug path string true "Slug of the SSP"
// @Param Body body endpointModel.GetTemplateReq true "endpointModel.GetTemplateReq"
// @Produce json
// @Success 200 {object} openrtb.BidResponse
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /rtb/{sspSlug} [post]
func (e *endpoint) getTemplate(ctx context.Context, r *http.Request) (any, error) {

	var endpointReq endpointModel.GetTemplateReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &endpointReq, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Получаем идентификатор SSP из URL
	sspSlug := r.PathValue("sspSlug")

	// Преобразуем запрос в бизнес-модель
	req := endpointReq.ConvertToBusinessModel(sspSlug)

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetTemplate(ctx, req)
}
