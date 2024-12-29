package beforeAnalytic

import "pkg/uuid"

type makeMappingImpressionIDs struct {
	baseLink
	exchangeRepository ExchangeRepository
}

func (r *makeMappingImpressionIDs) Apply(dto *beforeAnalytic) error {

	for _, impression := range dto.BidRequest.Impressions {
		dto.MappingExchangeImpressionIDs[impression.ID] = uuid.New()
	}

	return nil
}
