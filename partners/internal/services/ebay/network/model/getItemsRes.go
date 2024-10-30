package model

import "partners/internal/services/ebay/model"

type GetItemsRes struct {
}

func (r GetItemsRes) ConvertToBusinessModel() []model.Item {
	return nil
}
