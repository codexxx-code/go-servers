package model

type GetADMReq struct {
	ID      string `validation:"required" schema:"-"`
	IsAdult bool   `validation:"required" schema:"is_adult"`
}
