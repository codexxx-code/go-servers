package model

type Prompt struct {
	ID   int    `db:"id" json:"id"`
	Case string `db:"case" json:"case"`
	Text string `db:"text" json:"text"`
}
