package http

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"exchange/internal/metrics"
	"exchange/internal/services/billing/model"
	"pkg/errors"
	"pkg/http/decoder"
	"pkg/validator"
)

// BillURL Засчитывает вызов NURL/BURL
// @Summary Биллит показ
// @Tags billing
// @Param Query query model.BillURLReq true "model.BillURLReq"
// @Success 200
// @Failure 400,500 {object} errors.Error
// @Router /billing/:id [get]
func (h endpoint) BillURL(ctx *fiber.Ctx) error {

	var req model.BillURLReq

	defer func() {
		metrics.IncBillingCallFromSSP(req.SSPSlug)
	}()

	// Парсим идентификатор запроса из пути
	req.ID = ctx.Params("id")

	// Парсим query параметры
	if err := decoder.DecodeFiber(ctx, decoder.DecodeSchema, &req); err != nil {
		return err
	}

	// Валидируем запрос
	if err := validator.Validate(req); err != nil {
		return err
	}

	// Вызываем функцию сервиса
	// Не передаем контекст запроса, так как нам не надо, чтобы при прерывании запроса прерывалась работа биллинга
	if errs := h.service.BillURL(context.Background(), req); len(errs) != 0 {
		var errTexts []any
		for i, err := range errs {
			errTexts = append(errTexts, fmt.Sprintf("error_%d", i+1), err.Error())
		}
		return errors.InternalServer.New("error while billing", errors.ParamsOption(errTexts...))
	}

	ctx.Status(fiber.StatusOK)

	return nil
}
