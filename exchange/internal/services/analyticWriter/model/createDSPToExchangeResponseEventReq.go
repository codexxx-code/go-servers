package model

import (
	"exchange/internal/services/analyticWriter/model/clickhouseTime"
)

type CreateDSPToExchangeResponseEventReq struct {
	ExchangeBidID        string                         `json:"exchange_bid_id"`        // Идентификатор бида
	ExchangeImpressionID string                         `json:"exchange_impression_id"` // Идентификатор импрешена
	ExchangeID           string                         `json:"exchange_id"`            // Идентификатор запроса
	RequestDateTime      clickhouseTime.ClickhouseTime  `json:"request_date_time"`      // Дата и время запроса
	SSPRequestImpression AnalyticRequestImpressionModel `json:"ssp_request_impression"` // Импрешен от SSP
	DSPRequestImpression AnalyticRequestImpressionModel `json:"dsp_request_impression"` // Импрешен в DSP
	DSPResponseBid       AnalyticResponseBidModel       `json:"dsp_response_bid"`       // Бид от DSP
}
