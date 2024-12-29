package beforeRequestToDSP

import (
	"pkg/decimal"
	"pkg/errors"
)

type dividePriceByClickunderDrumSize struct {
	baseLink
}

func (r *dividePriceByClickunderDrumSize) Apply(dto *beforeRequestToDSP) (err error) {

	// Если запрос не кликандер, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проходимся по каждому Impression
	for i, impression := range dto.BidRequest.Impressions {

		if dto.SSP.ClickunderDrumSize == nil {
			return errors.InternalServer.New("ClickunderDrumSize is nil",
				errors.ParamsOption("SSP", dto.SSP.Slug),
			)
		}

		// BidFloor = BidFloor / Размер барабана
		impression.BidFloor = impression.BidFloor.Div(decimal.NewFromInt(int(*dto.SSP.ClickunderDrumSize)))

		dto.BidRequest.Impressions[i] = impression
	}

	return nil
}
