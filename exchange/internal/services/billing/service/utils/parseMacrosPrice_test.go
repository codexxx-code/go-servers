package utils

import (
	"testing"

	"pkg/decimal"
	"pkg/errors"
	"pkg/testUtils"
)

func TestParseMacrosPrice(t *testing.T) {
	type args struct {
		macrosPrice string
	}
	tests := []struct {
		name    string
		args    args
		want    decimal.Decimal
		wantErr error
	}{
		{
			name: "1. Успешный парсинг правильно раскрытого макроса. Число с точкой",
			args: args{
				macrosPrice: "1.23",
			},
			want:    decimal.NewFromFloat(1.23),
			wantErr: nil,
		},
		{
			name: "2. Успешный парсинг правильно раскрытого макроса. Целое число",
			args: args{
				macrosPrice: "1",
			},
			want:    decimal.NewFromFloat(1),
			wantErr: nil,
		},
		{
			name: "3. Успешный парсинг неправильно раскрытого макроса. Число с запятой",
			args: args{
				macrosPrice: "1,23",
			},
			want:    decimal.NewFromFloat(1.23),
			wantErr: nil,
		},
		{
			name: "4. Успешный парсинг неправильно раскрытого макроса. Число с долларом",
			args: args{
				macrosPrice: "$1.23",
			},
			want:    decimal.NewFromFloat(1.23),
			wantErr: nil,
		},
		{
			name: "5. Ошибка при парсинге нераскрытого макроса",
			args: args{
				macrosPrice: "${AUCTION_PRICE}",
			},
			want:    decimal.NewFromFloat(0),
			wantErr: errors.BadRequest.New(""),
		},
		{
			name: "6. Ошибка при парсинге пустой строки",
			args: args{
				macrosPrice: "",
			},
			want:    decimal.NewFromFloat(0),
			wantErr: errors.BadRequest.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMacrosPrice(tt.args.macrosPrice)
			testUtils.CheckError(t, err, tt.wantErr, false)
			if !got.Equal(tt.want) {
				t.Errorf("ParseMacrosPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
