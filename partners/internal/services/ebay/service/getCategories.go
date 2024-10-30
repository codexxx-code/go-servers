package service

import (
	"context"

	"partners/internal/services/ebay/model"
	ebayNetworkModel "partners/internal/services/ebay/network/model"
)

func (s *EbayService) GetCategories(ctx context.Context, req model.GetCategoriesReq) ([]model.Category, error) {

	// Получаем ID дерева категорий для региона US
	categoryTreeIDRes, err := s.ebayNetwork.GetCategoryTreeID(ctx, ebayNetworkModel.GetCategoryTreeIDReq{
		MarketplaceID: "EBAY_US",
	})
	if err != nil {
		return nil, err
	}

	// Получаем категории из дерева
	categories, err := s.ebayNetwork.GetCategories(ctx, ebayNetworkModel.GetCategoriesReq{
		CategoryTreeID: categoryTreeIDRes.CategoryTreeID,
	})
	if err != nil {
		return nil, err
	}

	return categories.ConvertToBusinessModel(req.MaxDeepLevel), nil
}
