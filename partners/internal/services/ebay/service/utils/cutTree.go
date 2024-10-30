package utils

import "partners/internal/services/ebay/model"

func CutTree(categories []model.Category, deep uint8) []model.Category {

	if deep == 0 {
		return nil
	}

	for i, category := range categories {
		categories[i].Subcategories = CutTree(category.Subcategories, deep-1)
	}

	return categories
}
