package beforeResponseToSSP

type addADM struct {
	baseLink
}

func (r *addADM) Apply(dto *beforeResponseToSSP) error {

	// Проверяем, что ADM еще не установлен
	if dto.chainSettings.admAlreadySet {
		return nil
	}

	// Проставляем настройку, что ADM уже установлен
	defer func() {
		dto.chainSettings.admAlreadySet = true
	}()

	currentSeatBidIndex := 0

	// Проходимся по каждому победившему биду
	for _, wonBid := range dto.wonBids {

		// Просто присваиваем AdMarkup победившему биду
		dto.bidResponse.SeatBids[currentSeatBidIndex].Bids[0].AdMarkup = wonBid.BidResponse.SeatBids[wonBid.SeatBidIndex].Bids[wonBid.BidIndex].AdMarkup
		currentSeatBidIndex++
	}

	return nil
}
