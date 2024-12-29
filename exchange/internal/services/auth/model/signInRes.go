package model

type SignInRes struct {
	Tokens Tokens // Токены доступа
	ID     string // Идентификатор пользователя
}
