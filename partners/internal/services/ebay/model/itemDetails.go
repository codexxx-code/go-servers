package model

type ItemDetails struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       PriceModel `json:"price"`
	Images      []string   `json:"images"`
	Brand       string     `json:"brand"`
	ItemWebUrl  string     `json:"itemWebUrl"`
}
