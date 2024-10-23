package model

type Prompt struct {
	ID   int    `db:"id"`
	Case string `db:"case"`
	Text string `db:"text"`
}
