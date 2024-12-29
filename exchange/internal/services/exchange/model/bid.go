package model

import (
	"pkg/openrtb"
)

type BidPointer struct {
	ExchangeBidID        string              // Идентификатор бида, генерируем сами
	ExchangeImpressionID string              // Идентификатор импрешена, который был в запросе от SSP
	BidResponse          openrtb.BidResponse // Полный ответ
	SeatBidIndex         int                 // Индекс SeatBid, в котором находится победивший Bid
	BidIndex             int                 // Победивший Bid

	DSPSlug              string // Слаг DSP
	IsAuctionSecondPrice bool   // Индикатор, вторая ли цена для аукциона
}

func (s *BidPointer) GetBid() openrtb.Bid {
	return s.BidResponse.SeatBids[s.SeatBidIndex].Bids[s.BidIndex]
}
