package model

type RefreshTokensReq struct {
	RefreshToken string `validate:"required"`
}
