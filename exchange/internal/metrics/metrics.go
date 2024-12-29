package metrics

import (
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"

	"pkg/errors"
)

const (
	sspSlugLabel     = "ssp_slug"
	dspSlugLabel     = "dsp_slug"
	statusCodeLabel  = "status_code"
	pathLabel        = "path"
	eventNameLabel   = "event_name"
	publisherIDLabel = "publisher_id"
)

type metrics struct {
	responseTimeMetric                  *prometheus.HistogramVec
	responseCodeFromAdExToSSP           *prometheus.CounterVec
	responseCodesFromDSPToAdEx          *prometheus.CounterVec
	responseCodesFromDSPToAdExOnBilling *prometheus.CounterVec
	billingCallsFromSSP                 *prometheus.CounterVec
	timeouts                            *prometheus.CounterVec
	eventCalls                          *prometheus.CounterVec
}

func (m *metrics) register() error {

	if err := prometheus.Register(m.responseTimeMetric); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.responseCodeFromAdExToSSP); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.responseCodesFromDSPToAdEx); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.responseCodesFromDSPToAdExOnBilling); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.billingCallsFromSSP); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.timeouts); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	if err := prometheus.Register(m.eventCalls); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}

var metricsInitialized atomic.Bool
var globalMetrics *metrics

func Init(namespace string) error {

	if metricsInitialized.Load() {
		return nil
	}
	metricsInitialized.Store(true)

	globalMetrics = &metrics{
		// Метрика для измерения времени ответа
		responseTimeMetric: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace:                       namespace,
			Subsystem:                       "",
			Name:                            "http_response_time_seconds",
			Help:                            "Histogram of response time in seconds.",
			ConstLabels:                     nil,
			Buckets:                         nil,
			NativeHistogramBucketFactor:     0,
			NativeHistogramZeroThreshold:    0,
			NativeHistogramMaxBucketNumber:  0,
			NativeHistogramMinResetDuration: 0,
			NativeHistogramMaxZeroThreshold: 0,
		}, []string{pathLabel, statusCodeLabel}),

		// Метрика по кодам ответа от AdEx к SSP
		responseCodeFromAdExToSSP: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "response_code_from_adex_to_ssp",
			Help:        "Count responses, grouped by HTTP-codes and ssp slugs",
			ConstLabels: map[string]string{},
		},
			[]string{statusCodeLabel, sspSlugLabel},
		),

		// Метрика по кодам ответа от DSP к AdEx
		responseCodesFromDSPToAdEx: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "response_codes_from_dsp_to_adex",
			Help:        "Count responses, grouped by HTTP-codes, dsp and ssp slugs",
			ConstLabels: map[string]string{},
		},
			[]string{statusCodeLabel, dspSlugLabel, sspSlugLabel},
		),

		// Метрика по кодам ответа от DSP к AdEx на биллинге
		responseCodesFromDSPToAdExOnBilling: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "response_codes_from_dsp_to_adex_on_billing",
			Help:        "Count responses, grouped by HTTP-codes, dsp slugs on billing",
			ConstLabels: map[string]string{},
		},
			[]string{statusCodeLabel, dspSlugLabel},
		),

		// Метрика по биллинговым вызовам от SSP
		billingCallsFromSSP: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "billing_calls_from_ssp",
			Help:        "Count billing calls, grouped by ssp slugs",
			ConstLabels: map[string]string{},
		},
			[]string{sspSlugLabel},
		),

		// Метрика для отслеживания таймаутов
		timeouts: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "timeouts",
			Help:        "Count timeouts, grouped by ssp slugs",
			ConstLabels: map[string]string{},
		},
			[]string{sspSlugLabel},
		),

		// Метрика для отслеживания ивентов
		eventCalls: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   "",
			Name:        "event_calls",
			Help:        "Count events, grouped by ssp slugs and type",
			ConstLabels: map[string]string{},
		},
			[]string{eventNameLabel, sspSlugLabel, publisherIDLabel},
		),
	}

	// Регистрируем метрики
	return globalMetrics.register()
}
