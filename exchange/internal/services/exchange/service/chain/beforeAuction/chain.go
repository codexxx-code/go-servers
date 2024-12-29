package beforeAuction

import (
	"context"

	"exchange/internal/services/exchange/model"
	exchangeRepository "exchange/internal/services/exchange/repository"
	"pkg/chain"
)

type beforeAuction struct {
	*model.AuctionDTO
}

// baseLink - базовая структура для всех модификаторов, которая реализует метод Run() интерфейса chain.Link и является первым звеном цепи
type baseLink struct {
	chain.Link[*beforeAuction]
}

// ExchangeRepository - контракт функционала репозитория, который должен быть передан в конструктор для работы чейна
var _ ExchangeRepository = new(exchangeRepository.ExchangeRepository)

type ExchangeRepository interface {
	GetPublisherVisibility(ctx context.Context, publisherID string) (model.PublisherVisibility, error)
}

// RunChain собирает и запускает цепочку модификаторов
func RunChain(
	exchangeRepository ExchangeRepository,
	dto *model.AuctionDTO,
) error {

	innerDTO := &beforeAuction{
		AuctionDTO: dto,
	}

	return chain.SetArrange[*beforeAuction](
		new(setIsClickunder),               // Определяет, является ли запрос запросом на кликандер
		new(setTrafficType),                // Определяет, тип трафика
		new(blockBy),                       // Проверка запросов на наличие в блоклистах
		new(changeTestTo0),                 // Заменяет значение теста на 0
		new(clickunderMultiplyImpressions), // Устанавливает количество исходящих запросов в зависимости от размера барабана кликандера
		new(setRequestCurrency),            // Устанавливает валюту запроса
		new(setPublisherID),                // Устанавливает идентификатор паблишера
		new(changeRequestID),               // Меняет идентификатор запроса на ExchangeID
		new(changeImpressionIDs),           // Меняет идентификаторы импрешенов на наши
		//&checkPublisherVisibility{exchangeRepository: exchangeRepository}, // Проверяем видимость паблишера
	).Run(innerDTO)
}
