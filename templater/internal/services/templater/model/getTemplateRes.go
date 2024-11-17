package model

type GetTemplateRes struct {
	ID       uint32 `json:"id"`
	SSPSlug  string `json:"ssp_slug"`
	Template string `json:"template"`
}
