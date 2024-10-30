package service

import (
	"context"

	"partners/internal/services/ebay/model"
	"partners/internal/services/ebay/service/utils"
	"pkg/errors"
)

func (s *EbayService) GetBreadcrumbs(ctx context.Context, req model.GetBreadcrumbsReq) (res model.Category, err error) {

	// Получаем все категории
	categories, err := s.GetCategories(ctx, model.GetCategoriesReq{}) //nolint:exhaustruct
	if err != nil {
		return res, err
	}

	// Ищем родителя по ID дочернего элемента
	if parentCategory := utils.GetParentByID(categories, req.ChildID); parentCategory != nil {
		return *parentCategory, nil
	} else {
		return res, errors.NotFound.New("child category by id not found")
	}
}
