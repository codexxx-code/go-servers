package testutil_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"pkg/openrtb/testutil"
)

type mockT struct {
	*testing.T
	failed bool
}

func (m *mockT) Error(...interface{}) {
	m.failed = true
}

func (m *mockT) Errorf(string, ...interface{}) {
	m.failed = true
}

func TestErrorIs(t *testing.T) {
	testcases := []struct {
		err    error
		target error
		failed bool
	}{
		{
			err:    fmt.Errorf("wrapped error: %w", io.EOF),
			target: errors.New("error"),
			failed: true,
		},
		{
			err:    fmt.Errorf("wrapped error: %w", io.EOF),
			target: io.EOF,
			failed: false,
		},
		{
			target: io.EOF,
			failed: true,
		},
		{
			err:    io.EOF,
			failed: true,
		},
	}

	for _, tc := range testcases {
		mock := &mockT{T: t}
		testutil.ErrorIs(mock, tc.err, tc.target)
		assertBool(t, tc.failed, mock.failed)
	}
}

type jsonT struct {
	Field int `json:"field"`
}

func TestEqualJSON(t *testing.T) {
	testcases := []struct {
		data     []byte
		expected interface{}
		failed   bool
	}{
		{
			data:     []byte(`{"field":1}`),
			expected: &jsonT{1},
			failed:   false,
		},
		{
			data:     []byte(`{"field":1}`),
			expected: &jsonT{2},
			failed:   true,
		},
		{
			expected: jsonT{1},
			failed:   true,
		},
		{
			data:     []byte(`{]`),
			expected: &jsonT{1},
			failed:   true,
		},
	}

	for _, tc := range testcases {
		mock := &mockT{T: t}
		testutil.EqualJSON(mock, tc.data, tc.expected)
		assertBool(t, tc.failed, mock.failed)
	}
}

func assertBool(t *testing.T, expected, actual bool) {
	t.Helper()

	if expected != actual {
		t.Errorf("\nValues are not equal: expected %v, actual %v", expected, actual)
	}
}
