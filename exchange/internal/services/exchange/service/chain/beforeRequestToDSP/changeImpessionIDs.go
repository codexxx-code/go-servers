package beforeRequestToDSP

import "pkg/uuid"

type changeImpressionIDs struct {
	baseLink
}

func (r *changeImpressionIDs) Apply(dto *beforeRequestToDSP) error {

	// Проходимся по каждому импрешену
	for i, impression := range dto.BidRequest.Impressions {

		// Генерируем новый ID для импрешена
		newID := uuid.New()

		// Добавляем новый идентификатор в маппинг, чтобы потом быстро вернуть все к изначальным идентификаторам
		dto.MappingImpressionsByDSPs.Set(dto.dsp.Slug, newID, impression.ID)

		// Заменяем идентификатор импрешена на новый
		dto.BidRequest.Impressions[i].ID = newID
	}

	return nil
}
