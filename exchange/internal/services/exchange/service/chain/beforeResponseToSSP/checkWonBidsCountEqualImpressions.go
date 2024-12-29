package beforeResponseToSSP

import "pkg/errors"

type checkWonBidsCountEqualImpressions struct {
	baseLink
}

func (r *checkWonBidsCountEqualImpressions) Apply(dto *beforeResponseToSSP) (err error) {

	// Проверяем, что количество выигранных бидов равно количеству импрешенов
	if len(dto.wonBids) != len(dto.BidRequest.Impressions) {
		return errors.InternalServer.New("Number of won bids is not equal to number of impressions")
	}

	return nil
}
