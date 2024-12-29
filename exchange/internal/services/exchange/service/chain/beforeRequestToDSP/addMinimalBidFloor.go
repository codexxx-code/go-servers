package beforeRequestToDSP

import "pkg/decimal"

type addMinimalBidFloor struct {
	baseLink
}

const minimalBidFloorRUB = 10
const minimalBidFloorUSD = 0.01

func (r *addMinimalBidFloor) Apply(dto *beforeRequestToDSP) error {

	// Проходимся по каждому Impression
	for i := range dto.BidRequest.Impressions {

		if dto.BidRequest.Impressions[i].BidFloor.Equal(decimal.Zero) {
			if dto.BidRequest.Impressions[i].BidFloorCurrency == "RUB" {
				dto.BidRequest.Impressions[i].BidFloor = decimal.NewFromInt(minimalBidFloorRUB)
			}

			if dto.BidRequest.Impressions[i].BidFloorCurrency == "USD" {
				dto.BidRequest.Impressions[i].BidFloor = decimal.NewFromFloat(minimalBidFloorUSD)
			}
		}

	}

	return nil
}
