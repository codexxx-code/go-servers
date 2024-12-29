package model

import (
	"exchange/internal/services/analyticWriter/model/clickhouseTime"
)

type CreateDSPBillingEventReq struct {
	ExchangeBidID                 string                         `json:"exchange_bid_id"`                    // Идентификатор бида
	ExchangeImpressionID          string                         `json:"exchange_impression_id"`             // Идентификатор импрешена
	ExchangeID                    string                         `json:"exchange_id"`                        // Идентификатор запроса
	RequestDateTime               clickhouseTime.ClickhouseTime  `json:"request_date_time"`                  // Дата и время запроса
	SSPRequestImpression          AnalyticRequestImpressionModel `json:"ssp_request_impression"`             // Импрешен от SSP
	DSPRequestImpression          AnalyticRequestImpressionModel `json:"dsp_request_impression"`             // Импрешен в DSP
	DSPResponseBid                AnalyticResponseBidModel       `json:"dsp_response_bid"`                   // Бид от DSP
	SSPResponseBid                AnalyticResponseBidModel       `json:"ssp_response_bid"`                   // Бид в SSP
	FactPriceSSPInDefaultCurrency string                         `json:"fact_price_ssp_in_default_currency"` // Фактическая цена, по которой нас забиллила SSP
	FactPriceDSPInDefaultCurrency string                         `json:"fact_price_dsp_in_default_currency"` // Фактическая цена, по которой мы забиллили DSP
}
