package beforeAuction

type changeImpressionIDs struct {
	baseLink
}

func (r *changeImpressionIDs) Apply(dto *beforeAuction) error {

	for i, impression := range dto.BidRequest.Impressions {
		dto.BidRequest.Impressions[i].ID = dto.MappingExchangeImpressionIDs[impression.ID]
	}

	return nil
}
