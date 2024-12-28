package model

import "generator/internal/enum/promptCase"

type UpdatePromptTemplateReq struct {
	Case     promptCase.PromptCase `json:"case" enums:"createHoroscope" validate:"required"`
	Template *string               `json:"template"`
}
