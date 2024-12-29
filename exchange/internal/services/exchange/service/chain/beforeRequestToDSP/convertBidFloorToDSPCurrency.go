package beforeRequestToDSP

import "pkg/currencyConverter"

type convertBidFloorToDSPCurrency struct {
	baseLink
}

func (r *convertBidFloorToDSPCurrency) Apply(dto *beforeRequestToDSP) (err error) {

	// Присваиваем валюту DSP в массив доступных для торгов валют
	dto.BidRequest.Currencies = []string{string(dto.dsp.Currency)}

	// Проходимся по каждому Impression
	for i, impression := range dto.BidRequest.Impressions {

		// Конвертируем BidFloor в валюту DSP
		dto.BidRequest.Impressions[i].BidFloor, err = currencyConverter.Convert(
			impression.BidFloor,
			impression.BidFloorCurrency,
			string(dto.dsp.Currency),
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		// Присваиваем валюту DSP в BidFloorCurrency
		dto.BidRequest.Impressions[i].BidFloorCurrency = string(dto.dsp.Currency)
	}

	return nil
}
