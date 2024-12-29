package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	analyticWriterSerivce "exchange/internal/services/analyticWriter/service"
	currencyService "exchange/internal/services/currency/service"
	dspModel "exchange/internal/services/dsp/model"
	dspService "exchange/internal/services/dsp/service"
	eventService "exchange/internal/services/event/service"
	"exchange/internal/services/exchange/model"
	exchangeNetwork "exchange/internal/services/exchange/network"
	exchangeRepository "exchange/internal/services/exchange/repository"
	fraudScoreModel "exchange/internal/services/fraudScore/model"
	fraudScoreService "exchange/internal/services/fraudScore/service"
	settingsModel "exchange/internal/services/setting/model"
	settingService "exchange/internal/services/setting/service"
	sspModel "exchange/internal/services/ssp/model"
	sspService "exchange/internal/services/ssp/service"
	"pkg/decimal"
	"pkg/openrtb"
)

// Запись для проверки реализации контрактов
var _ ExchangeRepository = new(exchangeRepository.ExchangeRepository)

type ExchangeRepository interface {
	CreateDSPResponse(context.Context, int, model.DSPResponse) error
	GetDSPResponses(context.Context, model.GetDSPResponsesReq) (bidResponses []model.DSPResponse, err error)

	GetCountryByIP(ipStr string) (countryCode string, err error)
	GetPublisherVisibility(ctx context.Context, publisherID string) (model.PublisherVisibility, error)

	SaveADM(ctx context.Context, bidID string, adm string) error
	GetADM(ctx context.Context, bidID string) (string, error)

	SaveUnusedBids(ctx context.Context, bids []model.BidPointer) error
	PopUnusedBids(ctx context.Context, count int) (bids []model.BidPointer, err error)

	SaveAnalyticDTOs(ctx context.Context, analyticDTOs []model.AnalyticDTO) error
}

var _ ExchangeNetwork = new(exchangeNetwork.ExchangeNetwork)

type ExchangeNetwork interface {
	SendBidRequestToDSP(context.Context, int, string, openrtb.BidRequest) (openrtb.BidResponse, int, error)
}

var _ SSPService = new(sspService.SSPService)

type SSPService interface {
	GetSSPs(context.Context, sspModel.GetSSPsReq) ([]sspModel.SSP, error)
}

var _ DSPService = new(dspService.DSPService)

type DSPService interface {
	GetDSPs(context.Context, dspModel.GetDSPsReq) ([]dspModel.DSP, error)
}

var _ CurrencyService = new(currencyService.CurrencyService)

type CurrencyService interface {
	GetRates(ctx context.Context) (map[string]decimal.Decimal, error)
}

var _ SettingService = new(settingService.SettingService)

type SettingService interface {
	GetSettings(context.Context) (settingsModel.Settings, error)
}

var _ EventService = new(eventService.EventService)

type EventService interface {
	CreateSSPBidRequestToExchangeEvent(context.Context, openrtb.BidRequest) error
	CreateExchangeBidResponseToSSPEvent(context.Context, openrtb.BidResponse) error
	CreateExchangeBidRequestToDSPEvent(context.Context, int, openrtb.BidRequest) error
}

var _ FraudScoreService = new(fraudScoreService.FraudScoreService)

type FraudScoreService interface {
	IsFraud(ctx context.Context, req fraudScoreModel.IsFraudReq) (bool, error)
}

var _ AnalyticWriterService = new(analyticWriterSerivce.AnalyticWriterService)

type AnalyticWriterService interface {
	CreateSSPToExchangeRequestEvent(context.Context, analyticWriterModel.CreateSSPToExchangeRequestEventReq) error
	CreateExchangeToDSPRequestEvent(context.Context, analyticWriterModel.CreateExchangeToDSPRequestEventReq) error
	CreateDSPToExchangeResponseEvent(context.Context, analyticWriterModel.CreateDSPToExchangeResponseEventReq) error
	CreateExchangeToSSPResponseEvent(context.Context, analyticWriterModel.CreateExchangeToSSPResponseEventReq) error
}

type ExchangeService struct {
	exchangeRepository    ExchangeRepository
	exchangeNetwork       ExchangeNetwork
	sspService            SSPService
	dspService            DSPService
	currencyService       CurrencyService
	settingsService       SettingService
	eventService          EventService
	fraudScoreService     FraudScoreService
	analyticWriterService AnalyticWriterService
}

func NewExchangeService(
	sspRepository ExchangeRepository,
	sspService SSPService,
	dspService DSPService,
	currencyService CurrencyService,
	settingsService SettingService,
	exchangeNetwork ExchangeNetwork,
	eventService EventService,
	fraudScoreService FraudScoreService,
	analyticWriterService AnalyticWriterService,
) *ExchangeService {
	return &ExchangeService{
		exchangeRepository:    sspRepository,
		sspService:            sspService,
		dspService:            dspService,
		currencyService:       currencyService,
		settingsService:       settingsService,
		exchangeNetwork:       exchangeNetwork,
		eventService:          eventService,
		fraudScoreService:     fraudScoreService,
		analyticWriterService: analyticWriterService,
	}
}
