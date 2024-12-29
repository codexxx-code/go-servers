package service

import (
	"context"
	"time"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/errors"
	"pkg/log"
)

func (s *ExchangeService) createAnalyticEventRequestFromSSP(ctx context.Context, dto *model.AuctionDTO) error {

	// Собираем сжатую информацию о запросе
	var (
		domain, bundle        string
		country, region, city string
		requestDatetime       = time.Now()
	)

	if dto.BidRequest.App != nil {
		bundle = dto.BidRequest.App.Bundle
	}

	if dto.BidRequest.Site != nil {
		domain = dto.BidRequest.Site.Domain
	}

	if dto.BidRequest.User != nil && dto.BidRequest.User.Geo != nil {
		country = dto.BidRequest.User.Geo.Country
		region = dto.BidRequest.User.Geo.Region
		city = dto.BidRequest.User.Geo.City
	}

	if dto.BidRequest.Device != nil && dto.BidRequest.Device.Geo != nil {
		if country != "" {
			country = dto.BidRequest.Device.Geo.Country
		}
		if region != "" {
			region = dto.BidRequest.Device.Geo.Region
		}
		if city != "" {
			city = dto.BidRequest.Device.Geo.City
		}
	}

	for _, impression := range dto.BidRequest.Impressions {

		var sspRequestImpression analyticWriterModel.AnalyticRequestImpressionModel

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

		// Конвертируем бидфлур в дефолтную валюту системы
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
		sspRequestImpression = analyticWriterModel.AnalyticRequestImpressionModel{
			RequestID:    dto.BidRequest.ID,
			ImpressionID: impression.ID,
			Slug:         dto.SSP.Slug,
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

		// Получаем наш внутренний идентификатор импрешена из маппинга
		exchangeImpressionID, ok := dto.MappingExchangeImpressionIDs[impression.ID]
		if !ok {
			return errors.InternalServer.New("Не смогли найти идентификатор импрешена в маппинге")
		}

		// Добавляем данные в DTO аналитики по импрешенам
		analyticDTOForImpression := dto.AnalyticDTOByImpression[exchangeImpressionID]

		analyticDTOForImpression.RequestDateTime = requestDatetime
		analyticDTOForImpression.SSPRequestImpression = sspRequestImpression
		analyticDTOForImpression.ExchangeImpressionID = dto.MappingExchangeImpressionIDs[impression.ID]
		analyticDTOForImpression.ExchangeID = dto.ExchangeID

		dto.AnalyticDTOByImpression[exchangeImpressionID] = analyticDTOForImpression

		// Сохраняем информацию о запросе в аналитику
		go func() {
			if err := s.analyticWriterService.CreateSSPToExchangeRequestEvent(ctx, analyticDTOForImpression.ConvertToSSPRequest()); err != nil {
				log.Error(ctx, err)
			}
		}()
	}

	return nil
}
