package network

import (
	"context"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"

	"pkg/errors"
)

func (s *BillingNetwork) BillDSP(
	_ context.Context,
	url string,
) (statusCode int, err error) {

	// Создаем request
	req := fasthttp.AcquireRequest()

	// Устанавливаем URL
	req.SetRequestURI(url)

	// Устанавливаем таймаут
	req.SetTimeout(5 * time.Second)

	// Устанавливаем метод
	req.Header.SetMethod(http.MethodGet)

	// Определяем response
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// Делаем запрос
	if err := s.client.Do(req, resp); err != nil {
		switch {
		case errors.Is(err, fasthttp.ErrTimeout):
			return 0, errors.Timeout.Wrap(err)
		default:
			return 0, errors.InternalServer.Wrap(err)
		}
	}

	// Проверяем код ответа
	statusCode = resp.StatusCode()
	switch {
	case 200 <= statusCode && statusCode <= 299:
		return statusCode, nil
	default:
		return statusCode, errors.InternalServer.New("HTTP code is not 2xx", errors.ParamsOption(
			"code", statusCode,
		))
	}
}
