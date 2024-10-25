package model

type DeletePromptReq struct {
	ID uint32 `json:"id" schema:"id" validate:"required"`
}
