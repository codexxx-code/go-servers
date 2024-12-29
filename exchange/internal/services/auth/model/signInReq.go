package model

type SignInReq struct {
	Login    string `validate:"required"` // Электронная почта пользователя
	Password string `validate:"required"` // Пароль пользователя
}
