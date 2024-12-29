package beforeRequestToDSP

import (
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/chain"
	"pkg/openrtb"
)

// baseLink - базовая структура для всех модификаторов, которая реализует функцию Run интерфейса chain.Link
type baseLink struct {
	chain.Link[*beforeRequestToDSP]
}

type beforeRequestToDSP struct {
	*model.AuctionDTO
	BidRequest openrtb.BidRequest
	dsp        dspModel.DSP
}

// RunChain собирает и запускает цепочку модификаторов
func RunChain(
	dto *model.AuctionDTO,
	bidRequest openrtb.BidRequest,
	dsp dspModel.DSP,
) (openrtb.BidRequest, error) {

	// Делаем для этой горутины с чейнами полную копию исходного запроса, чтобы не испортить его. Так как у нас внутри хранятся
	// слайсы, они будут меняться по ссылке, даже если передаем объект копированием
	// А еще у нас эта функция выполняется ассинхронно и мы словим такую проблему, что в рандомный момент времени разные
	// чейны разных горутин будут пытаться изменить один и тот же объект (вспоминаем базу, ведь у нас слайс имеет ссылку
	// на исходный массив под собой, и в каждой горутине эта ссылка на один и тот же объект памяти), что приведет,
	// во первых, к гонкам данных, во вторых, к непредсказуемому состоянию объекта в момент отправки запроса в DSP
	innerDTO := &beforeRequestToDSP{
		AuctionDTO: dto,
		BidRequest: bidRequest.Copy(),
		dsp:        dsp,
	}

	return innerDTO.BidRequest, chain.SetArrange[*beforeRequestToDSP](
		new(reduceTMax),                      // изменяет максимальное время ответа на коэффициент, заданный в настройках системы
		new(clickunder),                      // изменяет запрос, в случае, если запрашивается контент типа clickunder
		new(addMinimalBidFloor),              // ставит минимальное значение bidFloor для DSP, если он пустой
		new(addOurMargin),                    // добавляет нашу маржу к BidFloor
		new(convertBidFloorToDSPCurrency),    // конвертирует BidFloor в валюту DSP
		new(applyAuctionType),                // определяем тип аукциона для DSP
		new(dividePriceByClickunderDrumSize), // устанавливает ставку согласно размеру барабана
		new(changeUserID),                    // изменяет ID пользователя
		new(changeRequestID),                 // изменяет ID запроса на рандомный
		new(changeImpressionIDs),             // изменяет ID импрешенов на рандомные
	).Run(innerDTO)
}
