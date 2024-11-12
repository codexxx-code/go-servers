package service

import (
	"context"

	"partners/internal/services/ebay/model"
	networkModel "partners/internal/services/ebay/network/model"
)

func (s *EbayService) GetItemsSummary(ctx context.Context, req model.GetItemsSummaryReq) ([]model.ItemSummary, error) {

	// Получаем все товары
	items, err := s.ebayNetwork.GetItemsSummary(ctx, networkModel.GetItemsSummaryReq{
		CategoryID: req.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	// Конвертируем их в наш вид и возвращаем
	return items.ConvertToBusinessModel(), nil
}
