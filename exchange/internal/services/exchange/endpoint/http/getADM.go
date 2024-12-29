package http

import (
	"github.com/gofiber/fiber/v2"

	"exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/validator"
)

// GetADM Возвращает ADM по идентификатору BidRequest
// @Summary Возвращает ADM по идентификатору BidRequest
// @Tags adm
// @Produce xml
// @Success 200
// @Failure 400,500 {object} errors.Error
// @Router /rtb/adm/{id} [get]
func (h *endpoint) GetADM(ctx *fiber.Ctx) (err error) {

	var req model.GetADMReq

	// Парсим идентификатор запроса из пути
	req.ID = ctx.Params("id")

	// Валидируем структуру
	if err = validator.Validate(req); err != nil {
		return err
	}

	// Вызываем функцию сервиса
	res, err := h.service.GetADM(ctx.Context(), req)
	if err != nil {
		return err
	}

	ctx.Set("Content-Type", "text/html; charset=utf-8")
	if _, err = ctx.Write([]byte(res)); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
