package beforeAuction

import (
	"pkg/slices"

	"exchange/internal/enum/formatType"
	"pkg/errors"
)

type setIsClickunder struct {
	baseLink
}

func (r *setIsClickunder) Apply(dto *beforeAuction) error {

	// Если есть хоть один тип запроса у первого импрешена, то ничего не делаем
	if dto.BidRequest.Impressions[0].AssetCount() != 0 {
		return nil
	}

	// Если SSP не поддерживает кликандер, но к нам пришел запрос на кликандер, то возвращаем ошибку
	if !slices.Contains(dto.SSP.FormatTypes, formatType.Clickunder) {
		return errors.BadRequest.New("SSP does not support clickunder traffic",
			errors.ParamsOption("sspSlug", dto.SSP.Slug))
	}

	// Определяем, что запрос является кликандером
	dto.IsClickunder = true
	dto.DrumSize = dto.SSP.ClickunderDrumSize

	return nil
}
