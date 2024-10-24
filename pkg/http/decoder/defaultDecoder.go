package decoder

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"

	"pkg/datetime"
	"pkg/errors"
	"pkg/reflectUtils"
)

type DecodeMethod int

const (
	DecodeSchema DecodeMethod = iota + 1
	DecodeJSON
)

func Decoder(
	ctx context.Context,
	r *http.Request,
	dest any,
	decodeSchemas ...DecodeMethod,
) (err error) {

	// Проверяем типы данных
	if err = reflectUtils.CheckPointerToStruct(dest); err != nil {
		return err
	}

	// Проходимся по каждому
	for _, decodeSchema := range decodeSchemas {
		switch decodeSchema {
		case DecodeSchema:
			decoder := schema.NewDecoder()
			decoder.RegisterConverter(datetime.Date{}, dateConverter) //nolint:exhaustruct
			err = decoder.Decode(dest, r.URL.Query())
		case DecodeJSON:
			err = json.NewDecoder(r.Body).Decode(dest)
		default:
			break
		}
		if err != nil {
			return errors.BadRequest.Wrap(
				err,
				errors.SkipThisCallOption(),
			)
		}
	}

	return nil
}

var dateConverter = func(val string) reflect.Value {

	date, err := datetime.ParseDate(val)
	if err != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(date)
}
