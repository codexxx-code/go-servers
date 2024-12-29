package beforeAuction

import "pkg/errors"

type clickunderMultiplyImpressions struct {
	baseLink
}

func (r *clickunderMultiplyImpressions) Apply(dto *beforeAuction) error {

	// Если запрос не кликандер, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	if dto.SSP.ClickunderDrumSize == nil {
		return errors.InternalServer.New("ClickunderDrumSize is nil",
			errors.ParamsOption("SSP", dto.SSP.Slug),
		)
	}

	// Дублируем первый импрешен на один раз меньше (первый импрешен же уже есть), на сколько нам нужно наполнить барабан
	// Для дальнейшей логики мы будем работать с ним, будто к нам пришел мультиимпрешен запросом от SSP
	for range *dto.SSP.ClickunderDrumSize - 1 {
		dto.BidRequest.Impressions = append(dto.BidRequest.Impressions, dto.BidRequest.Impressions[0])
	}

	return nil
}
