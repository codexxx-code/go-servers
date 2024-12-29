package metrics

import (
	"context"

	"pkg/log"
)

func IncEventCall(eventName, ssp, publisherID string) {

	if globalMetrics.eventCalls == nil {
		log.Error(context.Background(), "eventCalls prometheus metric not initialized")
		return
	}

	globalMetrics.eventCalls.WithLabelValues(eventName, ssp, publisherID).Inc()
}
