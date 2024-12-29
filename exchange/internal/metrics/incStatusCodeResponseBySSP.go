package metrics

import (
	"context"
	"strconv"

	"pkg/log"
)

func IncStatusCodeBySSPMiddleware(sspSlug string, statusCode int) {
	if globalMetrics.responseCodeFromAdExToSSP == nil {
		log.Error(context.Background(), "responseCodeFromAdExToSSP prometheus metric not initialized")
		return
	}

	// Записываем информацию о времени ответа с использованием прометеуса
	globalMetrics.responseCodeFromAdExToSSP.WithLabelValues(
		strconv.Itoa(statusCode),
		preparePath(sspSlug),
	).Inc()
}
