package metrics

import (
	"context"

	"pkg/log"
)

func IncTimeouts(ssp string) {

	if globalMetrics.timeouts == nil {
		log.Error(context.Background(), "timeouts prometheus metric not initialized")
		return
	}

	globalMetrics.timeouts.WithLabelValues(ssp).Inc()
}
