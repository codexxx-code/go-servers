package testutil

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
)

// ErrorIs asserts that the target is included in the chain of errors.
func ErrorIs(tb testing.TB, err error, target error) {
	tb.Helper()

	if !errors.Is(err, target) {
		tb.Errorf(
			"\nTarget is not included in the chain of errors:\n\t- target:\n\t\t%s\n\t- chain:\n\t\t%s\n",
			printError(target),
			printErrors(err, "\n\t\t"),
		)
	}
}

func printError(err error) string {
	if err == nil {
		return "<nil>"
	}

	var builder strings.Builder
	builder.Grow(64) //nolint:gomnd

	writeError(&builder, err)

	return builder.String()
}

func printErrors(err error, indent string) string {
	if err == nil {
		return "<nil>"
	}

	var builder strings.Builder
	builder.Grow(128) //nolint:gomnd

	writeError(&builder, err)

	for err = errors.Unwrap(err); err != nil; {
		builder.WriteString(indent)
		writeError(&builder, err)
		err = errors.Unwrap(err)
	}

	return builder.String()
}

func writeError(builder *strings.Builder, err error) {
	builder.WriteByte('"')
	builder.WriteString(err.Error())
	builder.WriteByte('"')
}

// EqualJSON asserts that object equal JSON from data bytes.
func EqualJSON(tb testing.TB, data []byte, expected interface{}) {
	tb.Helper()

	rv := reflect.ValueOf(expected)
	if rv.Kind() != reflect.Ptr {
		tb.Error("\nExpected value must be a non-nil pointer to a structure.")
		return
	}

	var actual interface{} = expected

	rt := reflect.TypeOf(actual).Elem()
	actual = reflect.New(rt).Interface()

	err := json.Unmarshal(data, actual)
	if err != nil {
		tb.Error(err)
		return
	}

	EqualValues(tb, expected, actual)
}

// EqualValues asserts that two objects are equal.
func EqualValues(tb testing.TB, expected, actual interface{}) {
	tb.Helper()

	if !reflect.DeepEqual(expected, actual) {
		tb.Errorf(
			"\nValues are not equal:\n\t- expected:\n\t\t%#v\n\t- actual:\n\t\t%#v",
			expected,
			actual,
		)
	}
}
