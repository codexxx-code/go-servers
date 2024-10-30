package utils

import "partners/internal/services/ebay/model"

func GetByID(categories []model.Category, id string) *model.Category {

	// Проходимся по всем категориям
	for _, category := range categories {

		// Если эта категория имеет нужный ID, возвращаем ее
		if category.ID == id {
			return &category
		}

		// Ищем нужную категорию в подкатегориях
		if targetCategory := GetByID(category.Subcategories, id); targetCategory != nil {
			return targetCategory
		}
	}

	return nil
}
