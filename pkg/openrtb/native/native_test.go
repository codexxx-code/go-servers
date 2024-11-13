package native_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"pkg/openrtb/testutil"
)

var files *testutil.Files

func TestMain(m *testing.M) {
	dir := filepath.Join(testutil.Testdata(), "native")

	cached, err := testutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	files = cached

	code := m.Run()
	os.Exit(code)
}

func assertEqualJSON(t *testing.T, name string, expected interface{}) {
	data := files.Get(name + ".json")
	if len(data) == 0 {
		t.Fatalf("\nFile not found or empty: %s", name)
	}

	testutil.EqualJSON(t, data, expected)
}

type Validater interface {
	Validate() error
}

type Testcase struct {
	Name      string
	Validater Validater
	Err       error
}

func assertValidate(t *testing.T, testcases []Testcase) {
	t.Helper()

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			testutil.ErrorIs(t, tc.Validater.Validate(), tc.Err)
		})
	}
}
