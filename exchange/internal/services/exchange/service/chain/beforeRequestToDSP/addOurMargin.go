package beforeRequestToDSP

import "pkg/decimal"

type addOurMargin struct {
	baseLink
}

func (r *addOurMargin) Apply(dto *beforeRequestToDSP) error {

	// Проходимся по каждому Impression
	for i, impression := range dto.BidRequest.Impressions {

		// Добавляем нашу маржу // new bid floor = bid floor * (1 + ourMargin)
		dto.BidRequest.Impressions[i].BidFloor = impression.BidFloor.Mul(decimal.NewFromInt(1).Add(dto.Settings.Margin))
	}

	return nil
}
