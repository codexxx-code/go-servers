package service

import (
	"context"
	"strings"

	"exchange/internal/enum/billingType"
	"exchange/internal/metrics"
	dspModel "exchange/internal/services/dsp/model"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/openrtb"
)

func (s *BillingService) billDSP(ctx context.Context, dsp dspModel.DSP, dspResponse exchangeModel.DSPResponse) error {

	// Выбираем URL, по которому будем биллить DSP
	var billURL string
	switch dsp.BillingURLType {

	case billingType.BURL:
		billURL = dspResponse.
			BidResponse.
			SeatBids[dspResponse.SeatBidIndex].
			Bids[dspResponse.BidIndex].
			BillingURL

	case billingType.NURL:
		billURL = dspResponse.
			BidResponse.
			SeatBids[dspResponse.SeatBidIndex].
			Bids[dspResponse.BidIndex].
			NoticeURL
	default:
		return errors.BadRequest.New("Unavailable billing type")
	}

	// Раскрываем макрос с ценой DSP в валюте DSP
	billURL = strings.ReplaceAll(billURL, openrtb.AuctionPriceMacros, dspResponse.BillingPriceInDSPCurrency.String())

	// Биллим DSP
	statusCode, err := s.billingNetwork.BillDSP(ctx, billURL)
	metrics.IncStatusCodeFromDSPOnBilling(statusCode, dsp.Slug)
	if err != nil {
		return err
	}

	return nil
}
