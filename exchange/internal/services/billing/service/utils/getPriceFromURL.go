package utils

import (
	"context"

	"exchange/internal/services/billing/model"
	"pkg/decimal"
	"pkg/errors"
	"pkg/log"
	"pkg/openrtb"
)

func GetPriceFromURL(req model.BillURLReq) (price decimal.Decimal, err error) {
	return getPriceFromURL(req, false, false)
}

func getPriceFromURL(req model.BillURLReq, macrosPriceParsingFailed, hardcodedPriceParsingFailed bool) (price decimal.Decimal, err error) {

	switch {
	case req.MacrosPrice != openrtb.AuctionPriceMacros && !macrosPriceParsingFailed: // Если строка не является макросом и парсинг не провалился (в случае )

		// Пытаемся распарсить строку в число
		price, err = ParseMacrosPrice(req.MacrosPrice)
		if err != nil { // Если ошибка, логгируем
			log.Warning(context.Background(), errors.BadRequest.Wrap(err))
			macrosPriceParsingFailed = true
			break
		}

		// Если пустая цена, логгируем
		if price.IsZero() {
			log.Warning(context.Background(), errors.BadRequest.New("price is zero", errors.ParamsOption("price", req.MacrosPrice)))
			macrosPriceParsingFailed = true
			break
		}

		// Если парсинг прошел успешно, возвращаем цену
		return price, nil

	case req.HardcodedPrice != "" && !hardcodedPriceParsingFailed: // Если передано значение в запасном поле

		// Пытаемся распарсить значение
		price, err = decimal.NewFromString(req.HardcodedPrice)
		if err != nil {
			log.Warning(context.Background(), errors.BadRequest.Wrap(err, errors.ParamsOption("bid_price", req.HardcodedPrice)))
			hardcodedPriceParsingFailed = true
			break
		}

		// Если пустая цена, логгируем
		if price.IsZero() {
			log.Warning(context.Background(), errors.BadRequest.New("bid_price is zero", errors.ParamsOption("bid_price", req.HardcodedPrice)))
			hardcodedPriceParsingFailed = true
			break
		}

		// Если парсинг прошел успешно, возвращаем цену
		return price, nil

	default: // Если макрос не раскрыт и нет значения в запасном поле

		return price, errors.BadRequest.New("no price in request")
	}

	// Если мы здесь, то парсинг прошел неудачно, вызываем функцию снова
	return getPriceFromURL(req, macrosPriceParsingFailed, hardcodedPriceParsingFailed)
}
