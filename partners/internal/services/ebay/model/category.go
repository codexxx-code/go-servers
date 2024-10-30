package model

type Category struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Subcategories []Category `json:"subcategories,omitempty"`
}
