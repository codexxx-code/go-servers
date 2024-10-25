package model

type GetPromptsReq struct {
	Cases []string `json:"cases" schema:"cases"`
}
