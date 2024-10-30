package utils

import (
	"partners/internal/services/ebay/model"
)

func GetParentByID(categories []model.Category, id string) *model.Category {

	for _, category := range categories {
		if preparedCategory := hasID(category, id); preparedCategory != nil {
			return preparedCategory
		}
	}

	return nil
}

func hasID(category model.Category, id string) *model.Category {

	// Если ID категории совпадает с искомым
	if category.ID == id {

		// Очищаем все дочерние категории
		category.Subcategories = nil

		return &category
	}

	for _, subcategory := range category.Subcategories {
		if childCategory := hasID(subcategory, id); childCategory != nil {
			category.Subcategories = []model.Category{*childCategory}
			return &category
		}
	}

	return nil
}
