package decimal

import (
	"database/sql/driver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"

	"pkg/errors"
)

func (d Decimal) MarshalBSONValue() (bsontype.Type, []byte, error) {
	mongoDecimal128, err := primitive.ParseDecimal128(d.String())
	if err != nil {
		return bson.TypeDecimal128, nil, errors.InternalServer.Wrap(err)
	}
	return bson.MarshalValue(mongoDecimal128)
}

func (d *Decimal) UnmarshalBSONValue(t bsontype.Type, value []byte) error {

	if t != bson.TypeDecimal128 {
		return errors.InternalServer.New("invalid bson value type",
			errors.ParamsOption("type", t.String()),
		)
	}
	mongoDecimal128, _, ok := bsoncore.ReadDecimal128(value)
	if !ok {
		return errors.InternalServer.New("invalid bson string value")
	}

	decimal, err := NewFromString(mongoDecimal128.String())
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	*d = decimal

	return nil
}

func (d *Decimal) UnmarshalJSON(data []byte) (err error) {
	return d.Decimal.UnmarshalJSON(data)
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return d.Decimal.MarshalJSON()
}

func (d *Decimal) Scan(src interface{}) error {
	return d.Decimal.Scan(src)
}

func (d Decimal) Value() (driver.Value, error) {
	return d.Decimal.Value()
}
