package model

type GetBreadcrumbsReq struct {
	ChildID string `json:"childID" schema:"childID" validate:"required"`
}
