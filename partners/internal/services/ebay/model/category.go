package model

type Category struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Level         uint8      `json:"level"`
	Subcategories []Category `json:"subcategories,omitempty"`
}
