package beforeResponseToSSP

import (
	"pkg/decimal"
	"pkg/errors"
)

type reducePriceForOurMargin struct {
	baseLink
}

func (r *reducePriceForOurMargin) Apply(dto *beforeResponseToSSP) (err error) {

	// Если цену уже проставили, то ничего не делаем
	if !dto.chainSettings.priceAlreadySet {
		return errors.InternalServer.New("reducePriceForOurMargin вызван до проставления цены")
	}

	for seatBidIndex, seatBid := range dto.bidResponse.SeatBids {
		for bidIndex, bid := range seatBid.Bids {
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].Price = bid.Price.Mul(decimal.NewFromInt(1).Sub(dto.Settings.Margin)).Round() // Price = price * (1 - margin)
		}
	}

	return nil
}
