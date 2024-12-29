package model

import (
	dspModel "exchange/internal/services/dsp/model"
	"pkg/openrtb"
	"pkg/uuid"
)

type BidResponseFromDSP struct {
	DSP         dspModel.DSP        `json:"dsp"`
	BidResponse openrtb.BidResponse `json:"bidResponse"`
}

// ExtractBids возвращает указатели на все биды этого ответа
func (s *BidResponseFromDSP) ExtractBids(mappingImpressionsIDsByDSP MappingImpressionsByDSPsMu) []BidPointer {

	var bids []BidPointer

	// Проходимся по каждому seatBid в ответе
	for seatBidIndex, seatBid := range s.BidResponse.SeatBids {

		// Проходимся по каждому bid в seatBid
		for bidIndex, bid := range seatBid.Bids {

			// Получаем exchangeImpressionID для этого импрешена
			exchangeImpressionID, ok := mappingImpressionsIDsByDSP.Get(s.DSP.Slug, bid.ImpID)
			if !ok {
				continue
			}

			bids = append(bids, BidPointer{
				ExchangeBidID:        uuid.New(),
				ExchangeImpressionID: exchangeImpressionID,
				BidResponse:          s.BidResponse,
				SeatBidIndex:         seatBidIndex,
				BidIndex:             bidIndex,
				DSPSlug:              s.DSP.Slug,
				IsAuctionSecondPrice: s.DSP.AuctionSecondPrice,
			})

		}
	}

	return bids
}
