package beforeAuction

type changeRequestID struct {
	baseLink
}

func (r *changeRequestID) Apply(dto *beforeAuction) error {

	dto.BidRequest.ID = dto.ExchangeID

	return nil
}
