package model

type CreatePromptReq struct {
	Case     string `json:"case" validate:"required"`
	Language string `json:"language" validate:"required"`
	Text     string `json:"text" validate:"required"`
}
