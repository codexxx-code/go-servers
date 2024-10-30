package endpoint

import (
	"context"
	"net/http"

	_ "partners/internal/services/ebay/model"
	"pkg/errors"
)

// @Summary Получение товаров по фильтрам
// @Tags ebay
// @Param Query query model.GetItemsReq false "model.GetItemsReq"
// @Produce json
// @Success 200 {object} []model.Item
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /ebay/items [get]
func (e *endpoint) getItems(ctx context.Context, r *http.Request) (any, error) {

	return nil, errors.InternalServer.New("Это заглушка для доки")

	/*
		var req model.GetItemsReq

		// Декодируем запрос
		if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
			return nil, err
		}

		// Валидируем запрос
		if err := validator.Validate(req); err != nil {
			return nil, err
		}

		// Вызываем метод сервиса
		return e.service.GetItems(ctx, req)
	*/
}
