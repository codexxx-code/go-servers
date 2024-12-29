package beforeAnalytic

import (
	"pkg/errors"
	"pkg/openrtb"
)

type checkGeoByIP struct {
	baseLink
	exchangeRepository ExchangeRepository
}

func (r *checkGeoByIP) Apply(dto *beforeAnalytic) error {

	// Проверяем наличие объекта device
	if dto.BidRequest.Device == nil || dto.BidRequest.Device.IP == "" {
		return errors.BadRequest.New("device object or device.ip not passed in request")
	}

	if dto.BidRequest.User == nil {
		return errors.BadRequest.New("user object not passed in request")
	}

	// Получаем страну по IP адресу
	country, err := r.exchangeRepository.GetCountryByIP(dto.BidRequest.Device.IP)
	if err != nil {
		return err
	}

	// Переписываем Гео у входящего запроса
	geo := &openrtb.Geo{ //nolint:exhaustruct
		Country: country,
	}
	dto.BidRequest.Device.Geo = geo
	dto.BidRequest.User.Geo = geo

	dto.GeoCountry = country

	return nil
}
