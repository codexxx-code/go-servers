package service

import (
	"context"
	"fmt"

	"exchange/internal/config"
	"exchange/internal/services/ssp/model"
)

// Функция для построения URL
func (s *SSPService) buildEndpointURL(sspSlug string) string {
	return fmt.Sprintf("https://bid.%s/rtb/%s", config.Load().Host, sspSlug)
}

func (s *SSPService) FindSSPs(ctx context.Context, req model.FindSSPsReq) (res model.FindSSPsRes, err error) {
	// Получаем список SSP
	ssps, err := s.sspRepository.FindSSPs(ctx, req)
	if err != nil {
		return res, err
	}

	// Проходимся по каждому элементу списка и добавляем EndpointURL
	for i, ssp := range ssps {
		ssps[i].EndpointURL = s.buildEndpointURL(ssp.Slug)
	}

	// Получаем общее количество SSP
	sspsCount, err := s.sspRepository.GetSSPsCount(ctx, req)
	if err != nil {
		return res, err
	}

	return model.FindSSPsRes{
		SSPs:      ssps,
		SSPsCount: sspsCount,
	}, nil
}
