package beforeRequestToDSP

import "pkg/uuid"

type changeUserID struct {
	baseLink
}

func (r *changeUserID) Apply(dto *beforeRequestToDSP) error {

	if dto.BidRequest.User == nil {
		return nil
	}

	dto.BidRequest.User.ID = uuid.New()

	return nil
}
