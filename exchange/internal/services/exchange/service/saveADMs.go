package service

import (
	"context"

	"exchange/internal/services/exchange/model"
)

func (s *ExchangeService) saveADMs(ctx context.Context, bids []model.WonBid) error {

	// Проходимся по каждому победившему биду
	for _, wonBid := range bids {

		// Сохраняем ad markup в бд
		if err := s.exchangeRepository.SaveADM(ctx, wonBid.ExchangeBidID, wonBid.BidResponse.SeatBids[wonBid.SeatBidIndex].Bids[wonBid.BidIndex].AdMarkup); err != nil {
			return err
		}
	}

	return nil
}
