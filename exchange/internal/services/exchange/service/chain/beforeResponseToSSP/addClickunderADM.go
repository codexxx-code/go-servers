package beforeResponseToSSP

import (
	"strings"

	"exchange/internal/enum"
	"pkg/errors"
	"pkg/url"
)

type addClickunderADM struct {
	baseLink
}

func (r *addClickunderADM) Apply(dto *beforeResponseToSSP) error {

	// Проверяем, что ADM еще не установлен
	if dto.chainSettings.admAlreadySet {
		return nil
	}

	// Если мы торгуем с SSP не кликандер трафиком, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проставляем настройку, что ADM уже установлен
	defer func() {
		dto.chainSettings.admAlreadySet = true
	}()

	// Собираем все идентификаторы победивших бидов
	wonBidIDs := make([]string, 0, len(dto.wonBids))
	for _, wonBid := range dto.wonBids {
		wonBidIDs = append(wonBidIDs, wonBid.ExchangeBidID)
	}

	// Проверяем, что у SSP определен тип ADM для кликандера
	if dto.SSP.ClickunderADMFormat == nil {
		return errors.InternalServer.New("ClickunderADMType is nil",
			errors.ParamsOption("SSP", dto.SSP.Slug),
		)
	}

	// Проходимся по каждому биду из ответа
	for seatBidIndex, seatBid := range dto.bidResponse.SeatBids {
		for bidIndex := range seatBid.Bids {

			var (
				showcaseURL  string
				pageLanguage string
			)
			if dto.GeoCountry == "USA" {
				showcaseURL = "https://webview.games/"
				pageLanguage = "en"
			} else {
				showcaseURL = dto.Settings.ShowcaseURL
				pageLanguage = "ru"
			}

			// Формируем URL для получения ADM страницы сайта
			urlADM, err := url.BuildURL(
				showcaseURL,
				"/game/happy-bucket/play", // TODO: Получать URL динамически и забирать из кэша
				map[string]string{
					"adm_id":     strings.Join(wonBidIDs, ","),
					"is_adult":   "true",
					"ad_type":    "video", // TODO: Получать тип рекламы динамически
					"lang":       pageLanguage,
					"request_id": dto.ExchangeID,
				},
				nil,
				true,
			)
			if err != nil {
				return err
			}

			// Раскрываем макрос со ссылкой на витрину
			adm := strings.ReplaceAll(*dto.SSP.ClickunderADMFormat, enum.ADMURLMacros, urlADM)
			dto.bidResponse.SeatBids[seatBidIndex].Bids[bidIndex].AdMarkup = adm

		}
	}
	return nil
}
