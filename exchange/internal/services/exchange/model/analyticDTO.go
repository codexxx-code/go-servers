package model

import (
	"time"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/analyticWriter/model/clickhouseTime"
)

type AnalyticDTO struct {
	ExchangeBidID                 string                                             `bson:"_id"`                    // Идентификатор бида
	ExchangeImpressionID          string                                             `bson:"exchange_impression_id"` // Идентификатор импрешена
	ExchangeID                    string                                             `bson:"exchange_id"`            // Идентификатор запроса
	RequestDateTime               time.Time                                          `bson:"request_date_time"`      // Дата и время запроса
	SSPRequestImpression          analyticWriterModel.AnalyticRequestImpressionModel `bson:"ssp_request_impression"` // Импрешен от SSP
	DSPRequestImpression          analyticWriterModel.AnalyticRequestImpressionModel `bson:"dsp_request_impression"` // Импрешен в DSP
	DSPResponseBid                analyticWriterModel.AnalyticResponseBidModel       `bson:"dsp_response_bid"`       // Бид от DSP
	SSPResponseBid                analyticWriterModel.AnalyticResponseBidModel       `bson:"ssp_response_bid"`       // Бид в SSP
	FactPriceSSPInDefaultCurrency string                                             `bson:"-"`                      // Фактическая цена, по которой нас забиллила SSP
	FactPriceDSPInDefaultCurrency string                                             `bson:"-"`                      // Фактическая цена, по которой мы забиллили DSP
}

func (a *AnalyticDTO) ConvertToSSPRequest() analyticWriterModel.CreateSSPToExchangeRequestEventReq {
	return analyticWriterModel.CreateSSPToExchangeRequestEventReq{
		ExchangeImpressionID: a.ExchangeImpressionID,
		ExchangeID:           a.ExchangeID,
		RequestDateTime:      clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression: a.SSPRequestImpression,
	}
}

func (a *AnalyticDTO) ConvertToDSPRequest() analyticWriterModel.CreateExchangeToDSPRequestEventReq {
	return analyticWriterModel.CreateExchangeToDSPRequestEventReq{
		ExchangeImpressionID: a.ExchangeImpressionID,
		ExchangeID:           a.ExchangeID,
		RequestDateTime:      clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression: a.SSPRequestImpression,
		DSPRequestImpression: a.DSPRequestImpression,
	}
}

func (a *AnalyticDTO) ConvertToDSPResponse() analyticWriterModel.CreateDSPToExchangeResponseEventReq {
	return analyticWriterModel.CreateDSPToExchangeResponseEventReq{
		ExchangeBidID:        a.ExchangeBidID,
		ExchangeImpressionID: a.ExchangeImpressionID,
		ExchangeID:           a.ExchangeID,
		RequestDateTime:      clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression: a.SSPRequestImpression,
		DSPRequestImpression: a.DSPRequestImpression,
		DSPResponseBid:       a.DSPResponseBid,
	}
}

func (a *AnalyticDTO) ConvertToSSPResponse() analyticWriterModel.CreateExchangeToSSPResponseEventReq {
	return analyticWriterModel.CreateExchangeToSSPResponseEventReq{
		ExchangeBidID:        a.ExchangeBidID,
		ExchangeImpressionID: a.ExchangeImpressionID,
		ExchangeID:           a.ExchangeID,
		RequestDateTime:      clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression: a.SSPRequestImpression,
		DSPRequestImpression: a.DSPRequestImpression,
		DSPResponseBid:       a.DSPResponseBid,
		SSPResponseBid:       a.SSPResponseBid,
	}
}

func (a *AnalyticDTO) ConvertToSSPBilling() analyticWriterModel.CreateSSPBillingEventReq {
	return analyticWriterModel.CreateSSPBillingEventReq{
		ExchangeBidID:                 a.ExchangeBidID,
		ExchangeImpressionID:          a.ExchangeImpressionID,
		ExchangeID:                    a.ExchangeID,
		RequestDateTime:               clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression:          a.SSPRequestImpression,
		DSPRequestImpression:          a.DSPRequestImpression,
		DSPResponseBid:                a.DSPResponseBid,
		SSPResponseBid:                a.SSPResponseBid,
		FactPriceSSPInDefaultCurrency: a.FactPriceSSPInDefaultCurrency,
	}
}

func (a *AnalyticDTO) ConvertToDSPBilling() analyticWriterModel.CreateDSPBillingEventReq {
	return analyticWriterModel.CreateDSPBillingEventReq{
		ExchangeBidID:                 a.ExchangeBidID,
		ExchangeImpressionID:          a.ExchangeImpressionID,
		ExchangeID:                    a.ExchangeID,
		RequestDateTime:               clickhouseTime.ClickhouseTime{Time: a.RequestDateTime},
		SSPRequestImpression:          a.SSPRequestImpression,
		DSPRequestImpression:          a.DSPRequestImpression,
		DSPResponseBid:                a.DSPResponseBid,
		SSPResponseBid:                a.SSPResponseBid,
		FactPriceSSPInDefaultCurrency: a.FactPriceSSPInDefaultCurrency,
		FactPriceDSPInDefaultCurrency: a.FactPriceDSPInDefaultCurrency,
	}
}
