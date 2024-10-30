package model

type GetCategoriesReq struct {
	MaxDeepLevel *uint8 `json:"maxDeepLevel" schema:"maxDeepLevel"`
}
