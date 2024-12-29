package model

import "exchange/internal/services/analyticWriter/model/clickhouseTime"

type CreateExchangeToDSPRequestEventReq struct {
	ExchangeImpressionID string                         `json:"exchange_impression_id"` // Идентификатор импрешена
	ExchangeID           string                         `json:"exchange_id"`            // Идентификатор запроса
	RequestDateTime      clickhouseTime.ClickhouseTime  `json:"request_date_time"`      // Дата и время запроса
	SSPRequestImpression AnalyticRequestImpressionModel `json:"ssp_request_impression"` // Импрешен от SSP
	DSPRequestImpression AnalyticRequestImpressionModel `json:"dsp_request_impression"` // Импрешен в DSP
}
