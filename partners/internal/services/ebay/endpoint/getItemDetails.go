package endpoint

import (
	"context"
	"net/http"

	"partners/internal/services/ebay/model"
	_ "partners/internal/services/ebay/model"
	"pkg/validator"
)

// @Summary Получение подробной информации по товару
// @Tags ebay
// @Param Query query model.GetItemDetailsReq false "model.GetItemDetailsReq"
// @Produce json
// @Success 200 {object} model.ItemDetails
// @Failure 400,401,403,404,500 {object} errors.Error
// @Router /ebay/item/:id [get]
func (e *endpoint) getItemDetails(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetItemDetailsReq

	req.ID = r.PathValue("id")

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return e.service.GetItemDetails(ctx, req)
}
