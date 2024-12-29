package metrics

import (
	"context"
	"strconv"

	"pkg/log"
)

func IncStatusCodeFromDSPRTB(statusCode int, dsp, ssp string) {

	if globalMetrics.responseCodesFromDSPToAdEx == nil {
		log.Error(context.Background(), "responseCodesFromDSPToAdEx prometheus metric not initialized")
		return
	}

	globalMetrics.responseCodesFromDSPToAdEx.WithLabelValues(strconv.Itoa(statusCode), dsp, ssp).Inc()
}
