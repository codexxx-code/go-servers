package utils

import (
	"context"
	"strconv"

	"exchange/internal/services/exchange/network"
	"pkg/errors"
	"pkg/log"
)

func LogNetworkErrors(errs []error) {

	// Собираем логи по категориям
	var (
		errorsTimeout            int
		errors204                int
		errorsUnexpectedHTTPCode int
		errorsUnknown            []error
	)

	// Проходимся по каждой ошибке
	for _, err := range errs {
		switch {
		case errors.IsContextError(err):
			errorsTimeout++
		case errors.Is(err, network.ErrStatusCode204):
			errors204++
		case errors.Is(err, network.ErrUnexpectedHTTPCode):
			errorsUnexpectedHTTPCode++
		default:
			errorsUnknown = append(errorsUnknown, err)
		}
	}

	// Логируем ошибки
	log.Debug(context.Background(), "Errors from DSPs",
		log.ParamsOption(
			"Timeout", strconv.Itoa(errorsTimeout),
			"204", strconv.Itoa(errors204),
			"UnexpectedHTTPCode", strconv.Itoa(errorsUnexpectedHTTPCode),
		),
	)

	for _, err := range errorsUnknown {
		log.LogError(context.Background(), err)
	}
}
