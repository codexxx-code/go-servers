package utils

import (
	"encoding/json"

	"pkg/errors"
	"pkg/maps"
	"pkg/openrtb"
)

// FillExtField донасыщает bidRequest данными
func FillExtField(req openrtb.BidRequest, addExtValues map[string]any) (res openrtb.BidRequest, err error) {

	extValues := make(map[string]any)

	// Парсим существующие данные из Ext поля запроса
	if req.Ext != nil {
		if err = json.Unmarshal(req.Ext, &extValues); err != nil {
			return res, errors.InternalServer.Wrap(err)
		}
	}

	extValues = maps.Join(extValues, addExtValues)

	// Заливаем данные обратно в Ext
	if req.Ext, err = json.Marshal(extValues); err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	return req, nil
}
