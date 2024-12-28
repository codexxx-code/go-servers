package model

type PromptTemplate struct {
	Case     string `db:"case" json:"case"`
	Template string `db:"template" json:"template"`
}
