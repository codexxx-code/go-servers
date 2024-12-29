package beforeResponseToSSP

import (
	"exchange/internal/services/exchange/model/idType"
	"pkg/openrtb"
	"pkg/url"
)

type addNURLForClickunder struct {
	baseLink
}

func (r *addNURLForClickunder) Apply(dto *beforeResponseToSSP) error {

	// Если у нас уже установлены NURL, то ничего не делаем
	if dto.chainSettings.nurlAlreadySet {
		return nil
	}

	// Если у нас запрос не на кликандер, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проставляем настройку, что NURL уже установлен
	defer func() {
		dto.chainSettings.nurlAlreadySet = true
	}()

	// Формируем NURL
	nurl, err := url.BuildURL(
		dto.Settings.Host,            // Хост
		billingPathWithInterpolation, // Путь
		map[string]string{ // GET-параметры
			urlTypeParamField: nurlParamValue,
			priceParamField:   openrtb.AuctionPriceMacros,
			idTypeParamField:  string(idType.ExchangeID),
			sspSlugField:      dto.SSP.Slug,
		},
		map[string]string{ // Интерполированные параметры в пути
			interpolatedIDParam: dto.ExchangeID,
		},
		true,
	)
	if err != nil {
		return err
	}

	// Ставим NURL в первый бид для ответа SSP
	dto.bidResponse.SeatBids[0].Bids[0].NoticeURL = nurl

	return nil
}
