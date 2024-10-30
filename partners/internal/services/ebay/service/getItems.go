package service

import (
	"context"

	"partners/internal/services/ebay/model"
	networkModel "partners/internal/services/ebay/network/model"
)

func (s *EbayService) GetItems(ctx context.Context, req model.GetItemsReq) ([]model.Item, error) {

	// Получаем все товары
	items, err := s.ebayNetwork.GetItems(ctx, networkModel.GetItemsReq{
		CategoryID: req.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	// Конвертируем их в наш вид и возвращаем
	return items.ConvertToBusinessModel(), nil
}
