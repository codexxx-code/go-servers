package network

import (
	"bufio"
	"bytes"
	"context"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"exchange/internal/services/fraudScore/model"
	"pkg/errors"
	"pkg/url"
)

// CheckFraudScore проверяет на наличие нарушений через Fraudscore API.
func (fc *FraudScoreNetwork) CheckFraudScore(ctx context.Context, req model.IsFraudReq) (bool, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Формируем URL с параметрами.
	url, err := url.BuildURL(
		fc.url,
		"",
		map[string]string{
			"key":        fc.key,
			"event_type": req.Event,
			"at":         strconv.Itoa(int(req.AT)),
			"ua":         req.UserAgent,
			"ip":         req.IP,
		},
		nil,
		false,
	)
	if err != nil {
		return false, err
	}

	request.SetRequestURI(url)
	request.Header.SetMethod(http.MethodGet)

	// Смотрим, не закончился ли контекст
	if err = ctx.Err(); err != nil {
		return false, errors.Timeout.Wrap(err)
	}

	done := make(chan error, 1)

	// Делаем запрос
	go func() {
		done <- fc.client.Do(request, response)
	}()

	select {
	case <-ctx.Done(): // Если контекст закончился
		return false, errors.Timeout.Wrap(ctx.Err())
	case err = <-done: // Если запрос завершился
		// Смотрим на результат
		if err != nil {
			switch {
			case errors.Is(err, fasthttp.ErrTimeout):
				return false, errors.Timeout.Wrap(err)
			default:
				return false, errors.InternalServer.Wrap(err)
			}
		}
	}

	if response.StatusCode() != fasthttp.StatusOK {
		return false, errors.InternalServer.New("Status code is not 200")
	}

	// Анализируем ответ.
	scan := bufio.NewScanner(bytes.NewReader(response.Body()))
	for scan.Scan() {
		if len(scan.Text()) > 0 {
			return true, nil
		}
	}

	return false, nil
}
