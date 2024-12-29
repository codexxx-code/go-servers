package beforeAuction

import "pkg/errors"

type setRequestCurrency struct {
	baseLink
}

func (r *setRequestCurrency) Apply(dto *beforeAuction) error {

	// Проходимся по каждому импрешену
	for _, impression := range dto.BidRequest.Impressions {

		switch {
		case dto.RequestCurrency == "": // Если валюта запроса не указана

			// Берем устанавливаем валюту из bidFloor
			dto.RequestCurrency = impression.BidFloorCurrency

		case dto.RequestCurrency != impression.BidFloorCurrency: // Если валюта не совпадает с уже установленной
			return errors.BadRequest.New("RequestCurrency is not equal to impression.BidFloorCurrency",
				errors.ParamsOption(
					"firstCurrency", dto.RequestCurrency,
					"secondCurrency", impression.BidFloorCurrency,
				))
		}
	}

	return nil
}
