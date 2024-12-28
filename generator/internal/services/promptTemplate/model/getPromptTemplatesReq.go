package model

import "generator/internal/enum/promptCase"

type GetPromptTemplatesReq struct {
	Cases []promptCase.PromptCase `json:"cases" schema:"cases" enums:"createHoroscope"`
}
