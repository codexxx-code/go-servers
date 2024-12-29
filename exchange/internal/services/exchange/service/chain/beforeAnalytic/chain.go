package beforeAnalytic

import (
	"exchange/internal/services/exchange/model"
	exchangeRepository "exchange/internal/services/exchange/repository"
	"pkg/chain"
)

type beforeAnalytic struct {
	*model.AuctionDTO
}

// baseLink - базовая структура для всех модификаторов, которая реализует метод Run() интерфейса chain.Link и является первым звеном цепи
type baseLink struct {
	chain.Link[*beforeAnalytic]
}

// ExchangeRepository - контракт функционала репозитория, который должен быть передан в конструктор для работы чейна
var _ ExchangeRepository = new(exchangeRepository.ExchangeRepository)

type ExchangeRepository interface {
	GetCountryByIP(ipStr string) (countryCode string, err error)
}

// RunChain собирает и запускает цепочку модификаторов
func RunChain(
	exchangeRepository ExchangeRepository,
	dto *model.AuctionDTO,
) error {

	innerDTO := &beforeAnalytic{
		AuctionDTO: dto,
	}

	return chain.SetArrange[*beforeAnalytic](
		&checkGeoByIP{exchangeRepository: exchangeRepository}, // Определяет гео по IP и затирает его в запросе новыми данными
		new(makeMappingImpressionIDs),                         // Создает маппинг идентификаторов импрешенов
	).Run(innerDTO)
}
