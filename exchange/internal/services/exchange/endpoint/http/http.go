package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"

	"exchange/internal/config"
	"exchange/internal/metrics"
	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/service"
	fiberMiddleware "pkg/http/fiber"
	"pkg/openrtb"
)

type endpoint struct {
	service ExchangeService
}

var _ ExchangeService = new(service.ExchangeService)

type ExchangeService interface {
	BidSSP(context.Context, model.SSPBidReq) (openrtb.BidResponse, error)
	GetADM(context.Context, model.GetADMReq) (string, error)
}

func MountExchangeEndpoint(app *fiber.App, service ExchangeService) {
	app.Mount("/rtb", newExchangeEndpoint(service))
}

func newExchangeEndpoint(service ExchangeService) *fiber.App {

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

	app.Get("/adm/:id", e.GetADM)    // GET  /rtb/adm/:id
	app.All("/proxy", Proxy)         // ALL  /rtb/proxy
	app.Post("/:ssp_slug", e.BidSSP) // POST /rtb/:ssp_slug

	return app
}

func preparePath(path string) string {
	switch {
	case strings.Contains(path, "/adm"):
		return "/rtb/adm"
	default:
		return path
	}
}
