package beforeRequestToDSP

import "pkg/uuid"

type changeRequestID struct {
	baseLink
}

func (r *changeRequestID) Apply(dto *beforeRequestToDSP) error {

	dto.BidRequest.ID = uuid.New()

	return nil
}
