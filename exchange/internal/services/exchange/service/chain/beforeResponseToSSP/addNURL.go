package beforeResponseToSSP

import (
	"exchange/internal/services/exchange/model/idType"
	"pkg/openrtb"
	"pkg/url"
)

type addNURL struct {
	baseLink
}

const (
	billingPathWithInterpolation = "/billing/:id"
	urlTypeParamField            = "url_type"
	nurlParamValue               = "nurl"
	priceParamField              = "price"
	interpolatedIDParam          = "id"
	idTypeParamField             = "id_type"
	sspSlugField                 = "ssp_slug"
)

func (r *addNURL) Apply(dto *beforeResponseToSSP) error {

	// Если у нас уже установлены NURL, то ничего не делаем
	if dto.chainSettings.nurlAlreadySet {
		return nil
	}

	// Проставляем настройку, что NURL уже установлен
	defer func() {
		dto.chainSettings.nurlAlreadySet = true
	}()

	// Индикатор, какой seatBid мы сейчас заполняем в ответе
	currentSeatBidIndex := 0

	// Проходимся по каждому биду в ответе
	for _, wonBid := range dto.wonBids {

		// Формируем NURL
		nurl, err := url.BuildURL(
			dto.Settings.Host,            // Хост
			billingPathWithInterpolation, // Путь
			map[string]string{ // GET-параметры
				urlTypeParamField: nurlParamValue,
				priceParamField:   openrtb.AuctionPriceMacros,
				idTypeParamField:  string(idType.ExchangeBidID),
				sspSlugField:      dto.SSP.Slug,
			},
			map[string]string{ // Интерполированные параметры в пути
				interpolatedIDParam: wonBid.ExchangeBidID,
			},
			true,
		)
		if err != nil {
			return err
		}

		// Ставим NURL в нужный бид для ответа SSP
		dto.bidResponse.SeatBids[currentSeatBidIndex].Bids[0].NoticeURL = nurl

		// Инкрементируем индекс seatBid
		currentSeatBidIndex++
	}

	return nil
}
