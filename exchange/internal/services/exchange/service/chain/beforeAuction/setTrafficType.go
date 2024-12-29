package beforeAuction

import (
	"exchange/internal/enum/sourceTrafficType"
)

type setTrafficType struct {
	baseLink
}

func (r *setTrafficType) Apply(dto *beforeAuction) error {

	// Обозначаем дефолтное значение
	dto.SourceTrafficType = sourceTrafficType.Desktop

	switch {

	case dto.BidRequest.App != nil: // Если переданы данные по приложению
		dto.SourceTrafficType = sourceTrafficType.InApp

	case dto.BidRequest.Site != nil: // Если переданы данные по сайту
		dto.SourceTrafficType = sourceTrafficType.Desktop
	}

	// Если запрос на кликандер
	if dto.IsClickunder {

		// Если в UserAgent есть данные по WebView, то пришел запрос с мобилки
		//if utils.IsWebView(dto.BidRequest.Device.UserAgent) {
		//	dto.SourceTrafficType = sourceTrafficType.InApp
		//}

	}

	return nil
}
