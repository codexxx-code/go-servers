package beforeAuction

type changeTestTo0 struct {
	baseLink
}

func (r *changeTestTo0) Apply(dto *beforeAuction) error {

	dto.BidRequest.Test = 0

	return nil
}
