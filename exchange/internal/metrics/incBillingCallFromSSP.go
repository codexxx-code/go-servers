package metrics

import (
	"context"

	"pkg/log"
)

func IncBillingCallFromSSP(ssp string) {

	if globalMetrics.billingCallsFromSSP == nil {
		log.Error(context.Background(), "billingCallsFromSSP prometheus metric not initialized")
		return
	}

	globalMetrics.billingCallsFromSSP.WithLabelValues(ssp).Inc()
}
