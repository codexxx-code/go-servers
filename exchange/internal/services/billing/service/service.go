package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	analyticWriterService "exchange/internal/services/analyticWriter/service"
	billingNetwork "exchange/internal/services/billing/network"
	billingRepository "exchange/internal/services/billing/repository"
	dspModel "exchange/internal/services/dsp/model"
	dspService "exchange/internal/services/dsp/service"
	exchangeModel "exchange/internal/services/exchange/model"
	exchangeRepository "exchange/internal/services/exchange/repository"
)

type BillingService struct {
	dspService            DSPService
	billingRepository     BillingRepository
	exchangeRepository    ExchangeRepository
	billingNetwork        BillingNetwork
	analyticWriterService AnalyticWriterService
}

var _ BillingRepository = new(billingRepository.BillingRepository)

type BillingRepository interface {
	CreateDSPWinEvent(context.Context, billingRepository.CreateDSPEventWinReq) error
	CreateSSPWinEvent(context.Context, billingRepository.CreateSSPEventWinReq) error
}

var _ DSPService = new(dspService.DSPService)

type DSPService interface {
	GetDSPs(context.Context, dspModel.GetDSPsReq) ([]dspModel.DSP, error)
}

// TODO: Вынести в отдельный модуль
var _ ExchangeRepository = new(exchangeRepository.ExchangeRepository)

type ExchangeRepository interface {
	GetDSPResponses(context.Context, exchangeModel.GetDSPResponsesReq) (bidResponses []exchangeModel.DSPResponse, err error)

	GetAnalyticDTOs(context.Context, exchangeModel.GetAnalyticDTOsReq) (analyticDTOs []exchangeModel.AnalyticDTO, err error)
}

var _ BillingNetwork = new(billingNetwork.BillingNetwork)

type BillingNetwork interface {
	BillDSP(context.Context, string) (int, error)
}

var _ AnalyticWriterService = new(analyticWriterService.AnalyticWriterService)

type AnalyticWriterService interface {
	CreateSSPBillingEventReq(context.Context, analyticWriterModel.CreateSSPBillingEventReq) error
	CreateDSPBillingEventReq(context.Context, analyticWriterModel.CreateDSPBillingEventReq) error
}

func NewBillingService(
	dspService DSPService,
	billingRepository BillingRepository,
	exchangeRepository ExchangeRepository,
	billingNetwork BillingNetwork,
	analyticWriterService AnalyticWriterService,
) *BillingService {
	return &BillingService{
		dspService:            dspService,
		billingRepository:     billingRepository,
		exchangeRepository:    exchangeRepository,
		billingNetwork:        billingNetwork,
		analyticWriterService: analyticWriterService,
	}
}
