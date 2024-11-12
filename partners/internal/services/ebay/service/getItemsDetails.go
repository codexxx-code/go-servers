package service

import (
	"context"

	"partners/internal/services/ebay/model"
	networkModel "partners/internal/services/ebay/network/model"
)

func (s *EbayService) GetItemDetails(ctx context.Context, req model.GetItemDetailsReq) (res model.ItemDetails, err error) {

	// Получаем все товары
	items, err := s.ebayNetwork.GetItemDetails(ctx, networkModel.GetItemDetailsReq{
		ID: req.ID,
	})
	if err != nil {
		return res, err
	}

	// Конвертируем их в наш вид и возвращаем
	return items.ConvertToBusinessModel(), nil
}
