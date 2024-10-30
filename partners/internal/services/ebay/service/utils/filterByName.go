package utils

import (
	"strings"

	"partners/internal/services/ebay/model"
)

func FilterByName(categories []model.Category, name string) []model.Category {

	filteredCategories := make([]model.Category, 0)

	// Проходимся по всем категориям
	for _, category := range categories {

		// Если название категории содержит искомое название, добавляем ее в массив
		if strings.Contains(strings.ToLower(category.Name), strings.ToLower(name)) {
			filteredCategories = append(filteredCategories, model.Category{
				ID:            category.ID,
				Name:          category.Name,
				Level:         category.Level,
				Subcategories: nil,
			})
		}

		// Получаем все необходимые категории по подстроке, возвращаем со всеми их подкатегориями
		filteredCategories = append(filteredCategories, FilterByName(category.Subcategories, name)...)
	}

	return filteredCategories
}
