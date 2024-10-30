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

func (s GetCategoriesRes) ConvertToBusinessModel() []model.Category {
	return s.Root.convertToBusinessModel(0).Subcategories
}

func (s category) convertToBusinessModel(level uint8) model.Category {

	// Конвертируем категорию из формата ebay в наш формат
	category := model.Category{
		ID:            s.CategoryInfo.CategoryId,
		Name:          s.CategoryInfo.CategoryName,
		Level:         level,
		Subcategories: nil,
	}

	// Если есть подкатегории, аллоцируем память под них
	if len(s.Subcategories) != 0 {
		category.Subcategories = make([]model.Category, 0, len(s.Subcategories))
	}

	// Проходимся по каждой подкатегории
	for _, subcategory := range s.Subcategories {

		// Конвертируем подкатегорию из формата ebay в наш формат
		category.Subcategories = append(category.Subcategories, subcategory.convertToBusinessModel(level+1))
	}

	return category
}
