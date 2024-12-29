package utils

import (
	"regexp"
	"strings"

	"pkg/decimal"
	"pkg/errors"
)

var regexpCatchNonDigitAndDot = regexp.MustCompile(`[^.\d]`)

func ParseMacrosPrice(macrosPrice string) (decimal.Decimal, error) {

	// Если кто-то додумается вставить число с запятой, а не точкой
	// Меняем запятую на точку
	preparedMacrosPrice := strings.ReplaceAll(macrosPrice, ",", ".")

	// Заменяем все возможные символы, кроме точки и цифр на пустую строку
	// Потому что были инциденты, когда вставляли знак доллара перед ценой
	preparedMacrosPrice = regexpCatchNonDigitAndDot.ReplaceAllString(preparedMacrosPrice, "")

	// Пытаемся распарсить значение
	price, err := decimal.NewFromString(preparedMacrosPrice)
	if err != nil {
		return price, errors.BadRequest.Wrap(err, errors.ParamsOption("price", macrosPrice))
	}

	return price, nil
}
