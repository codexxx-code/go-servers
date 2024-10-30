package model

type UpdatePromptReq struct {
	ID       uint32  `json:"id" validate:"required"`
	Case     *string `json:"case"`
	Language *string `json:"language"`
	Text     *string `json:"text"`
}
