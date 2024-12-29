package service

import (
	"context"
	"exchange/internal/services/ssp/model"
	sspRepository "exchange/internal/services/ssp/repository"
	"exchange/internal/services/transactor"
)

type SSPService struct {
	sspRepository SSPRepository
	transactor    Transactor
}

func NewSSPService(
	sspRepository SSPRepository,
	transactor Transactor,
) *SSPService {
	return &SSPService{
		sspRepository: sspRepository,
		transactor:    transactor,
	}
}

var _ SSPRepository = new(sspRepository.SSPRepository)

type SSPRepository interface {
	CreateSSP(context.Context, model.CreateSSPReq) error
	DeleteSSP(context.Context, model.DeleteSSPReq) error
	UpdateSSP(context.Context, model.UpdateSSPReq) error
	FindSSPs(context.Context, model.FindSSPsReq) ([]model.SSP, error)
	GetSSPs(context.Context, model.GetSSPsReq) ([]model.SSP, error)
	GetSSPsCount(ctx context.Context, req model.FindSSPsReq) (count int, err error)
	LinkSSPToDSPs(ctx context.Context, req model.LinkSSPToDSPsReq) error
	UnlinkSSPToDSPs(ctx context.Context, req model.UnlinkSSPToDSPsReq) error
}

var _ Transactor = new(transactor.Transactor)

type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) error) error
}
