package model

import "exchange/internal/services/analyticWriter/model/clickhouseTime"

type CreateSSPToExchangeRequestEventReq struct {
	ExchangeImpressionID string                         `json:"exchange_impression_id"` // Идентификатор импрешена. Уникальный идентификатор записи
	ExchangeID           string                         `json:"exchange_id"`            // Идентификатор запроса
	RequestDateTime      clickhouseTime.ClickhouseTime  `json:"request_date_time"`      // Дата и время запроса
	SSPRequestImpression AnalyticRequestImpressionModel `json:"ssp_request_impression"` // Импрешен от SSP
}
