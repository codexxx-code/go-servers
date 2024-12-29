package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"exchange/internal/config"
	"exchange/internal/metrics"
	endpointModel "exchange/internal/services/exchange/endpoint/http/model"
	"pkg/errors"
	"pkg/http/decoder"
	"pkg/validator"
)

// BidSSP Проводит торги для запроса от SSP
// @Summary ORTB торги
// @Tags rtb
// @Accept json
// @Param Body body endpointModel.SSPBidReq true "endpointModel.SSPBidReq"
// @Success 204
// @Failure 400,500 {object} errors.Error
// @Router /rtb/{ssp_slug} [post]
func (h *endpoint) BidSSP(ctx *fiber.Ctx) error {

	sspSlug := utils.CopyString(ctx.Params("ssp_slug"))

	defer func() {
		metrics.IncStatusCodeBySSPMiddleware(sspSlug, ctx.Response().StatusCode())
	}()

	var req endpointModel.SSPBidReq

	// При любой ситуации, кроме успешной, возвращаем 204
	ctx.Status(fiber.StatusNoContent)

	// Декодируем тело запроса
	err := decoder.DecodeFiber(ctx, decoder.DecodeJSON, &req)
	if err != nil {
		return err
	}

	// Валидируем структуру
	if err = validator.Validate(req); err != nil {
		return err
	}

	// Запускаем функцию сервиса с аукционом и остальным интересным
	res, err := h.service.BidSSP(ctx.Context(), req.ConvertToBusinessModel(sspSlug))
	if err != nil {
		return err
	}

	// Смотрим на переменную, отвечающую за реакцию на успешный аукцион
	if config.Load().IsSendNoContentForSuccessResponse {
		return nil
	}

	// Отправляем RTB-ответ и 200 код
	if err = ctx.Status(fiber.StatusOK).JSON(res); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
