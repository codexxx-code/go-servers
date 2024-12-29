package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"

	"exchange/internal/metrics"

	"exchange/internal/config"
	"exchange/internal/services/billing/model"
	"exchange/internal/services/billing/service"
	fiberMiddleware "pkg/http/fiber"
)

type endpoint struct {
	service BillingService
}

var _ BillingService = new(service.BillingService)

type BillingService interface {
	BillURL(context.Context, model.BillURLReq) []error
}

func MountBillingEndpoint(app *fiber.App, service BillingService) {
	app.Mount("/billing", newBillingEndpoint(service))
}

func newBillingEndpoint(service BillingService) *fiber.App {

	e := &endpoint{
		service: service,
	}

	var app *fiber.App

	// Смотрим на переменную, отвечающую за реакцию на ошибки
	if config.Load().IsSendNoContentInsteadErrorsInResponse { // Если на любые ошибки шлем 204, ставим кастомный хэндлер
		app = fiber.New(fiber.Config{ //nolint:exhaustruct
			ErrorHandler: fiberMiddleware.ErrorHandler204Code,
		})
	} else { // Если отвечаем с ошибками (используем только для дева), ставим обычный REST хэндлер
		app = fiber.New(fiber.Config{ //nolint:exhaustruct
			ErrorHandler: fiberMiddleware.DefaultErrorHandler,
		})
	}

	app.Use(metrics.ResponseTimeMiddleware(preparePath))

	app.Get("/:id", e.BillURL) // GET /billing/:id

	return app
}

func preparePath(path string) string {
	switch {
	case strings.Contains(path, "/billing"):
		return "/billing"
	default:
		return path
	}
}
