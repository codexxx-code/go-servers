package beforeAuction

type setPublisherID struct {
	baseLink
}

func (r *setPublisherID) Apply(dto *beforeAuction) error {

	switch {
	case dto.BidRequest.Site != nil:
		dto.PublisherID = dto.BidRequest.Site.ID
	case dto.BidRequest.App != nil:
		dto.PublisherID = dto.BidRequest.Site.ID
	}

	return nil
}
