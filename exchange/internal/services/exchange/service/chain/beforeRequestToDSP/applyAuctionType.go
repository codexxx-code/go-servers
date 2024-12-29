package beforeRequestToDSP

type applyAuctionType struct {
	baseLink
}

func (r *applyAuctionType) Apply(dto *beforeRequestToDSP) error {

	if dto.dsp.AuctionSecondPrice {
		dto.BidRequest.AuctionType = 2
	} else {
		dto.BidRequest.AuctionType = 1
	}

	return nil
}
