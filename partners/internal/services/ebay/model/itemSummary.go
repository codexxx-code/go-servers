package model

type ItemSummary struct {
	ID             string          `json:"id"`             // Идентификатор
	Title          string          `json:"title"`          // Название
	Images         []string        `json:"images"`         // Изображения
	Price          PriceModel      `json:"price"`          // Текущая цена
	MarketingPrice *MarketingPrice `json:"marketingPrice"` // Дополнительные данные о цене товара, если есть скидка
	ItemWebURL     string          `json:"itemWebURL"`     // Ссылка на товар
}

type MarketingPrice struct {
	OriginalPrice      PriceModel `json:"originalPrice"`      // Изначальная цена
	DiscountPercentage string     `json:"discountPercentage"` // Процент скидки
	DiscountAmount     PriceModel `json:"discountAmount"`     // Размер скидки
}

type PriceModel struct {
	Value    string `json:"value"`    // Значение
	Currency string `json:"currency"` // Валюта
}
