package beforeResponseToSSP

import (
	"context"

	"exchange/internal/services/exchange/model"
	fraudScoreModel "exchange/internal/services/fraudScore/model"
	"pkg/chain"
	"pkg/openrtb"
)

// baseLink - базовая структура для всех модификаторов, которая реализует функцию Run интерфейс chain.Link
type baseLink struct {
	chain.Link[*beforeResponseToSSP]
}

type chainSettings struct {
	bidsAlreadyInit bool // Флаг, что биды уже инициализированы
	admAlreadySet   bool // Флаг, что ADM уже установлен
	priceAlreadySet bool // Флаг, что цена уже установлена
	nurlAlreadySet  bool // Флаг, что NURL уже установлен
}

type beforeResponseToSSP struct {
	*model.AuctionDTO
	bidResponse   openrtb.BidResponse // Формируемый в процессе работы чейна ответ в SSP
	chainSettings chainSettings       // Флаги, конфигурирующие работу чейна
	wonBids       []model.WonBid      // Ответы от DSP, которые выиграли аукцион
}

type FraudScoreService interface {
	IsFraud(ctx context.Context, req fraudScoreModel.IsFraudReq) (bool, error)
}

// RunChain собирает и запускает цепочку модификаторов
func RunChain(
	dto *model.AuctionDTO,
	wonBids []model.WonBid,
	fraudScoreService FraudScoreService,
	exchangeService ExchangeService,
) (openrtb.BidResponse, error) {

	innerDTO := &beforeResponseToSSP{
		AuctionDTO: dto,
		wonBids:    wonBids,
		chainSettings: chainSettings{
			bidsAlreadyInit: false,
			admAlreadySet:   false,
			priceAlreadySet: false,
			nurlAlreadySet:  false,
		},
		bidResponse: openrtb.BidResponse{}, //nolint:exhaustruct
	}

	return innerDTO.bidResponse, chain.SetArrange[*beforeResponseToSSP](

		// Проверки перед началом работы
		new(checkWonBidsCountEqualImpressions), // Проверяем, что количество выигранных бидов равно количеству импрешенов в запросе

		// Инициализация ответа
		new(initResponse),         // Добавляем идентификатор запроса
		new(initBidsForClikunder), // Добавляем один бид для барабана
		new(initBids),             // Добавляем такое количество бидов, сколько было импрешенов

		&checkFraudScore{fraudScoreService: fraudScoreService}, // Проверка на фродность

		// ADM
		new(addClickunderADM), // Формируем свой ADM для кликандера
		new(addADM),           // Добавляем ADM из ответа DSP

		// BillingPriceInDSPCurrency
		&addPriceForClickunder{exchangeService: exchangeService}, // Добавляем цену для кликандера (барабан)
		&addPrice{exchangeService: exchangeService},              // Добавляем валюту из ответа DSP

		new(addOtherFields), // Временный костыль, чтобы мы хотя бы начали побеждать у SSP

		// NURL
		new(addNURLForClickunder), // Формируем свой NURL для кликандера
		new(addNURL),              // Формируем свой NURL
		new(addPriceInNURL),       // Добавляем цену в NURL для необходимых SSP

		// Проверки после работы
		new(checkMaxPrice),           // Проверяем, что цена не превышает 500 рублей
		new(checkMaxPriceClickunder), // Проверяем, что цена не превышает 1000 рублей для кликандера
		new(checkAllSets),            // Проверяем, что все конфигурируемые необходимые параметры проставлены
	).Run(innerDTO)
}
