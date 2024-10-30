package model

import "partners/internal/services/ebay/model"

type GetCategoriesRes struct {
	Root category `json:"rootCategoryNode"`
}

type category struct {
	CategoryInfo  categoryInfo `json:"category"`
	Subcategories []category   `json:"childCategoryTreeNodes"`
}

type categoryInfo struct {
	CategoryId   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

func (s GetCategoriesRes) ConvertToBusinessModel(maxDeepLevel *uint8) []model.Category {
	return s.Root.convertToBusinessModel(maxDeepLevel).Subcategories
}

func (s category) convertToBusinessModel(maxDeepLevel *uint8) model.Category {

	// Конвертируем категорию из формата ebay в наш формат
	category := model.Category{
		ID:            s.CategoryInfo.CategoryId,
		Name:          s.CategoryInfo.CategoryName,
		Subcategories: nil,
	}

	// Если уровень вложенности больше максимального, то прекращаем рекурсию
	if maxDeepLevel != nil && *maxDeepLevel == 0 {
		return category
	}

	// Если есть подкатегории, аллоцируем память под них
	if len(s.Subcategories) != 0 {
		category.Subcategories = make([]model.Category, 0, len(s.Subcategories))
	}

	// Уменьшаем максимальный уровень вложенности
	if maxDeepLevel != nil {
		maxDeepLevelValue := *maxDeepLevel
		maxDeepLevel = new(uint8)
		*maxDeepLevel = maxDeepLevelValue - 1
	}

	// Проходимся по каждой подкатегории
	for _, subcategory := range s.Subcategories {

		// Конвертируем подкатегорию из формата ebay в наш формат
		category.Subcategories = append(category.Subcategories, subcategory.convertToBusinessModel(maxDeepLevel))
	}

	return category
}
