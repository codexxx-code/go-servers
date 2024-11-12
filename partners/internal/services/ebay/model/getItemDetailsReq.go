package model

type GetItemDetailsReq struct {
	ID string `json:"-" schema:"-" validate:"required"`
}
