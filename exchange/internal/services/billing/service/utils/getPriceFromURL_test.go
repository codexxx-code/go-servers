package utils

import (
	"testing"

	"exchange/internal/services/billing/model"
	"pkg/decimal"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func TestGetPriceFromURL(t *testing.T) {
	type args struct {
		req model.BillURLReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrice decimal.Decimal
		wantErr   error
	}{
		{
			name: "1. Успешное получение цены из раскрытого макроса",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    "1.23",
					HardcodedPrice: "9.99",
				},
			},
			wantPrice: decimal.NewFromFloat(1.23),
			wantErr:   nil,
		},
		{
			name: "2. Успешное получение цены из запасного поля, если макрос не раскрыт",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    openrtb.AuctionPriceMacros,
					HardcodedPrice: "1.23",
				},
			},
			wantPrice: decimal.NewFromFloat(1.23),
			wantErr:   nil,
		},
		{
			name: "3. Успешное получение цены из криво раскрытого макроса",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    "$1,23",
					HardcodedPrice: "9.99",
				},
			},
			wantPrice: decimal.NewFromFloat(1.23),
			wantErr:   nil,
		},
		{
			name: "4. Успешное получение цены из запасного поля при цене макроса равной нулю",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    "0",
					HardcodedPrice: "1.23",
				},
			},
			wantPrice: decimal.NewFromFloat(1.23),
			wantErr:   nil,
		},
		{
			name: "5. Успешное получение цены из запасного поля при совсем криво распаршенном макросе",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    "1.2.34.5.", // К такому мы не готовы совсем
					HardcodedPrice: "1.23",
				},
			},
			wantPrice: decimal.NewFromFloat(1.23),
			wantErr:   nil,
		},
		{
			name: "6. Ошибка при отсутствии запасного поля и раскрытого макроса",
			args: args{
				req: model.BillURLReq{
					MacrosPrice:    openrtb.AuctionPriceMacros,
					HardcodedPrice: "",
				},
			},
			wantPrice: decimal.NewFromFloat(0),
			wantErr:   errors.BadRequest.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrice, err := GetPriceFromURL(tt.args.req)
			testUtils.CheckError(t, err, tt.wantErr, false)
			if !gotPrice.Equal(tt.wantPrice) {
				t.Errorf("GetPriceFromURL() gotPrice = %v, want %v", gotPrice, tt.wantPrice)
			}
		})
	}
}
