package beforeAuction

import (
	"pkg/errors"
)

type blockBy struct {
	baseLink
}

func (r blockBy) Apply(dto *beforeAuction) error {

	if dto.BidRequest.Site != nil && inExistBlockList(dto.BidRequest.Site.Domain) {
		return errors.BadRequest.New("Block by site")
	}

	// TODO тут не просто блокировать, а потом распределять трафик на разные витрины

	/*	if dto.BidRequest.Device != nil {
		if dto.BidRequest.Device.Geo != nil {
			if dto.BidRequest.Device.Geo.Country != "USA" {
				return errors.InternalServer.New("Block by GEO")
			}
		}
	}*/

	return nil
}

func inExistBlockList(site string) bool {
	sites := map[string]bool{}

	return sites[site]
}
