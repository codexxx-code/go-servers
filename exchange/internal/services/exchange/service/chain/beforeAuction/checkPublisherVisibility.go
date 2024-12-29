package beforeAuction

import "pkg/errors"

type checkPublisherVisibility struct {
	baseLink
	exchangeRepository ExchangeRepository
}

const targetVisibility float32 = 0.6

func (r *checkPublisherVisibility) Apply(dto *beforeAuction) error {

	// Получаем все показы и загрузки по паблишеру
	visibilityObject, err := r.exchangeRepository.GetPublisherVisibility(dto.Ctx, dto.PublisherID)
	if err != nil {
		return err
	}

	// Проверяем, чтобы в статистике зачислилось хотя бы 100 загрузок, а то сразу забаним всех
	if visibilityObject.Loads < 100 {
		return nil
	}

	// Если visibility ниже целевой видимости, то отвергаем запрос
	publisherVisibility := visibilityObject.GetVisibility()
	if publisherVisibility < targetVisibility {
		return errors.BadRequest.New("У паблишера низкая видимость",
			errors.ParamsOption(
				"publisherID", dto.PublisherID,
				"targetVisibility", targetVisibility,
				"publisherVisibility", publisherVisibility,
			),
		)
	}

	return nil
}
