package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/errors"
	"pkg/openrtb"
)

func (s *ExchangeService) createAnalyticEventRequestToDSP(ctx context.Context, dto *model.AuctionDTO, bidRequest openrtb.BidRequest, dsp dspModel.DSP) error {

	// Собираем сжатую информацию о запросе в DSP
	var (
		domain, bundle        string
		country, region, city string
	)

	if bidRequest.App != nil {
		bundle = bidRequest.App.Bundle
	}

	if bidRequest.Site != nil {
		domain = bidRequest.Site.Domain
	}

	if bidRequest.User != nil && bidRequest.User.Geo != nil {
		country = bidRequest.User.Geo.Country
		region = bidRequest.User.Geo.Region
		city = bidRequest.User.Geo.City
	}

	if bidRequest.Device != nil && bidRequest.Device.Geo != nil {
		if country != "" {
			country = bidRequest.Device.Geo.Country
		}
		if region != "" {
			region = bidRequest.Device.Geo.Region
		}
		if city != "" {
			city = bidRequest.Device.Geo.City
		}
	}

	for _, impression := range bidRequest.Impressions {

		var dspRequestImpression analyticWriterModel.AnalyticRequestImpressionModel

		var (
			width, height int
			adType        string
		)

		switch {
		case impression.Banner != nil:
			width = impression.Banner.Width
			height = impression.Banner.Height
			adType = "banner"

		case impression.Video != nil:
			width = impression.Video.Width
			height = impression.Video.Height
			adType = "video"

		case impression.Audio != nil:
			adType = "audio"

		case impression.Native != nil:
			adType = "native"

		default:
			adType = "clickunder"
		}

		// Конвертируем бидфлур в дефолтную валюту системы для аналитики
		bidFloorInDefaultCurrency, err := currencyConverter.Convert(
			impression.BidFloor, // Бидфлур из запроса
			impression.BidFloorCurrency,
			currencyConverter.DefaultCurrency,
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		// Забираем информацию для аналитики из запроса
		dspRequestImpression = analyticWriterModel.AnalyticRequestImpressionModel{
			RequestID:    bidRequest.ID,
			ImpressionID: impression.ID,
			Slug:         dsp.Slug,
			Domain:       domain,
			Bundle:       bundle,
			Geo: analyticWriterModel.GeoModel{
				Country: country,
				Region:  region,
				City:    city,
			},
			Width:                     width,
			Height:                    height,
			AdType:                    adType,
			BidFloorInDefaultCurrency: bidFloorInDefaultCurrency.String(),
		}

		// Так как мы уже изменили у бид реквеста все идентификаторы импрешенов, то нам надо найти наш внутренний, для этого мы используем маппинг
		exchangeImpressionID, ok := dto.MappingImpressionsByDSPs.Get(dsp.Slug, impression.ID)
		if !ok {
			return errors.InternalServer.New("Не смогли найти идентификатор импрешена в маппинге")
		}

		// Получаем основную информацию по запросу из SSP
		analyticDTOForImpression := dto.AnalyticDTOByImpression[exchangeImpressionID]

		// Заполняем информацию о запросе в DSP
		analyticDTOForImpression.DSPRequestImpression = dspRequestImpression

		// Заполняем новую мапу в разрезе DSP - Импрешен
		dto.AnalyticDTOByDSPRequestByImpression.Set(dsp.Slug, impression.ID, analyticDTOForImpression)

		// Сохраняем информацию о запросе в DSP в аналитику
		if err = s.analyticWriterService.CreateExchangeToDSPRequestEvent(ctx, analyticDTOForImpression.ConvertToDSPRequest()); err != nil {
			return err
		}
	}

	return nil
}
