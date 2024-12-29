package model

import "pkg/decimal"

// AnalyticRequestImpressionModel - основная необходимая информация из импрешена запроса (к нам или от нас)
type AnalyticRequestImpressionModel struct {
	RequestID                 string   `json:"request_id" bson:"request_id"`                                       // Наш идентификатор запроса
	ImpressionID              string   `json:"impression_id" bson:"impression_id"`                                 // Идентификатор импрешена
	Slug                      string   `json:"slug" bson:"slug"`                                                   // Слаг объекта (dsp или ssp)
	Domain                    string   `json:"domain" bson:"domain"`                                               // Домен из запроса
	Bundle                    string   `json:"bundle" bson:"bundle"`                                               // Бандл из запроса
	Geo                       GeoModel `json:"geo" bson:"geo"`                                                     // Гео из запроса
	Width                     int      `json:"width" bson:"width"`                                                 // Ширина креатива из запроса
	Height                    int      `json:"height" bson:"height"`                                               // Высота креатива из запроса
	AdType                    string   `json:"ad_type" bson:"ad_type"`                                             // Тип рекламы из запроса
	BidFloorInDefaultCurrency string   `json:"bid_floor_in_default_currency" bson:"bid_floor_in_default_currency"` // Бидфлур из запроса
}

func GetClearAnalyticRequestImpressionModel() AnalyticRequestImpressionModel {
	return AnalyticRequestImpressionModel{
		RequestID:    "",
		ImpressionID: "",
		Slug:         "",
		Domain:       "",
		Bundle:       "",
		Geo: GeoModel{
			Country: "",
			Region:  "",
			City:    "",
		},
		Width:                     0,
		Height:                    0,
		AdType:                    "",
		BidFloorInDefaultCurrency: decimal.Zero.String(),
	}
}
