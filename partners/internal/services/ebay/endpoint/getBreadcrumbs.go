package endpoint

import (
	"context"
	"net/http"

	"partners/internal/services/ebay/model"
	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Получение всей цепочки категорий до указанной с самого верхнего уровня до указанной
// @Tags ebay
// @Param Query query model.GetBreadcrumbsReq false "model.GetBreadcrumbsReq"
// @Produce json
// @Success 200 {object} []model.Category
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /ebay/category [get]
func (e *endpoint) getBreadcrumbs(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetBreadcrumbsReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetBreadcrumbs(ctx, req)
}
