package model

type GetPromptsReq struct {
	Cases     []string `json:"cases" schema:"cases"`
	Languages []string `json:"languages" schema:"languages"`
}
