package service

import (
	"context"

	"exchange/internal/services/exchange/model"
)

func (s *ExchangeService) GetADM(ctx context.Context, req model.GetADMReq) (adm string, err error) {

	// Получаем ADM по идентификатору записи
	return s.exchangeRepository.GetADM(ctx, req.ID)
}
