package model

type DeleteDSPReq struct {
	Slug string `validate:"required"`
}
