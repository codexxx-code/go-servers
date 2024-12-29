package model

type DeleteSSPReq struct {
	Slug string `validate:"required"`
}
