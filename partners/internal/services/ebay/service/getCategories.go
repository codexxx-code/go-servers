package service

import (
	"context"

	"partners/internal/services/ebay/model"
	ebayNetworkModel "partners/internal/services/ebay/network/model"
	"partners/internal/services/ebay/service/utils"
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
	ebayCategories, err := s.ebayNetwork.GetCategories(ctx, ebayNetworkModel.GetCategoriesReq{
		CategoryTreeID: categoryTreeIDRes.CategoryTreeID,
	})
	if err != nil {
		return nil, err
	}

	// Конвертируем модель ebay в нашу
	categories := ebayCategories.ConvertToBusinessModel()

	// Если есть ID, фильтруем по нему
	if req.ID != nil {
		if targetCategory := utils.GetByID(categories, *req.ID); targetCategory != nil {
			categories = []model.Category{*targetCategory}
		}
	}

	// Если есть название, фильтруем по нему
	if req.Name != nil {
		categories = utils.FilterByName(categories, *req.Name)
	}

	// Если необходимо обрезать дерево, обрезаем
	if req.MaxDeepLevel != nil {
		categories = utils.CutTree(categories, *req.MaxDeepLevel)
	}

	return categories, nil
}
