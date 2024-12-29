package service

import (
	"sync"
	"time"

	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
)

func (s *ExchangeService) saveWonBids(dto *model.AuctionDTO, wonBids []model.WonBid) error {

	// Получаем коэффициент конвертации валюты из валюты SSP в базовую валюту. Когда умножаем на этот коэффициент, получаем цену в базовой валюте
	currencySSPCoefficient, err := currencyConverter.Coefficient(
		dto.RequestCurrency,
		currencyConverter.DefaultCurrency,
		dto.CurrencyRates,
	)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(wonBids)) // канал для передачи ошибок

	for bidIndex, wonBid := range wonBids {
		wg.Add(1) // увеличиваем счетчик горутин
		go func(bidIndex int, wonBid model.WonBid) {
			defer wg.Done() // уменьшаем счетчик по завершении горутины

			// Получаем коэффициент конвертации валюты
			currencyDSPCoefficient, err := currencyConverter.Coefficient(
				wonBid.BidResponse.Currency,
				currencyConverter.DefaultCurrency,
				dto.CurrencyRates,
			)
			if err != nil {
				errCh <- err // передаем ошибку в канал
				return
			}

			// Сохраняем оригинальный ответ от DSP
			err = s.exchangeRepository.CreateDSPResponse(dto.Ctx, bidIndex, model.DSPResponse{
				ExchangeID:                dto.ExchangeID,
				ExchangeBidID:             wonBid.ExchangeBidID,
				BidResponse:               wonBid.BidResponse,
				SeatBidIndex:              wonBid.SeatBidIndex,
				BidIndex:                  wonBid.BidIndex,
				BillingPriceInDSPCurrency: wonBid.BillingPriceInDSPCurrency,
				SlugDSP:                   wonBid.DSPSlug,
				RecordDateTime:            time.Time{}, // Заполняется в репозитории
				OriginRequestID:           dto.BidRequest.ID,
				CurrencySSPCoefficient:    currencySSPCoefficient,
				SSPCurrency:               dto.RequestCurrency,
				CurrencyDSPCoefficient:    currencyDSPCoefficient,
				DSPCurrency:               wonBid.BidResponse.Currency,
				SlugSSP:                   dto.SSP.Slug,
				RequestPublisherID:        dto.PublisherID,
				DrumSize:                  dto.DrumSize,
			})
			if err != nil {
				errCh <- err // передаем ошибку в канал
				return
			}
		}(bidIndex, wonBid) // передаем параметры внутрь горутины
	}

	wg.Wait()    // ждем завершения всех горутин
	close(errCh) // закрываем канал после завершения всех горутин

	// Проверяем, не возникло ли ошибок
	if len(errCh) > 0 {
		return <-errCh // возвращаем первую возникшую ошибку
	}
	return nil
}
