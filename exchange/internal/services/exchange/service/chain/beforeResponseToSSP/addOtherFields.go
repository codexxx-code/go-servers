package beforeResponseToSSP

type addOtherFields struct {
	baseLink
}

// Временный костыль, чтобы мы хотя бы начали побеждать у SSP
func (r *addOtherFields) Apply(dto *beforeResponseToSSP) error {

	for seatBidIndex, seatBid := range dto.bidResponse.SeatBids {
		for bidIndex, _ := range seatBid.Bids {
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].AdvDomains = []string{
				"vidimo.media",
			}
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].ImageURL = "https://cdn.adx.com.ru/banner/0000000000000000.jpg"

			// TODO думаю тут стоит делать такой костыль, только если эти значения реально пустые
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].CampaingID = "66c7182ba897d80001520631"
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].CreativeID = "66cd9992a897d800015d8af7"
		}
	}

	return nil
}
