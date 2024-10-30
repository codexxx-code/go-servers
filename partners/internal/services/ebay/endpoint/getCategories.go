package endpoint

import (
	"context"
	"net/http"

	"partners/internal/services/ebay/model"
	"pkg/http/decoder"
	"pkg/validator"
)

// @Summary Получение категорий товаров
// @Tags ebay
// @Param Query query model.GetCategoriesReq false "model.GetCategoriesReq"
// @Produce json
// @Success 200 {object} []model.Category
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /ebay/category [get]
func (e *endpoint) getCategories(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetCategoriesReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetCategories(ctx, req)
}
