package service

import (
	"context"

	"exchange/internal/services/dsp/model"
	dspRepository "exchange/internal/services/dsp/repository"
	"exchange/internal/services/transactor"
)

type DSPService struct {
	dspRepository DSPRepository
	transactor    Transactor
}

func NewDSPService(
	dspRepository DSPRepository,
	transactor Transactor,
) *DSPService {
	return &DSPService{
		dspRepository: dspRepository,
		transactor:    transactor,
	}
}

var _ DSPRepository = new(dspRepository.DSPRepository)

type DSPRepository interface {
	CreateDSP(context.Context, model.CreateDSPReq) error
	DeleteDSP(context.Context, model.DeleteDSPReq) error
	UpdateDSP(context.Context, model.UpdateDSPReq) error
	FindDSPs(context.Context, model.FindDSPsReq) ([]model.DSP, error)
	GetDSPs(context.Context, model.GetDSPsReq) ([]model.DSP, error)
	GetDSPsCount(context.Context, model.FindDSPsReq) (int, error)
}

var _ Transactor = new(transactor.Transactor)

type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}
