package metrics

import (
	"context"
	"strconv"

	"pkg/log"
)

func IncStatusCodeFromDSPOnBilling(statusCode int, dsp string) {

	if globalMetrics.responseCodesFromDSPToAdExOnBilling == nil {
		log.Error(context.Background(), "responseCodesFromDSPToAdEx prometheus metric not initialized")
		return
	}

	globalMetrics.responseCodesFromDSPToAdExOnBilling.WithLabelValues(strconv.Itoa(statusCode), dsp).Inc()
}
