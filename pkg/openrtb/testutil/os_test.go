package testutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"pkg/openrtb/testutil"
)

func TestTestdata(t *testing.T) {
	dir := testutil.Testdata()
	_, err := os.Stat(dir)
	assertBool(t, false, os.IsNotExist(err))
}

func TestReadDir(t *testing.T) {
	dir := t.TempDir()

	if _, err := os.MkdirTemp(dir, ""); err != nil {
		t.Fatal(err)
	}

	f, err := os.CreateTemp(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = f.Close() })

	if _, err := f.WriteString(f.Name()); err != nil {
		t.Fatal(err)
	}

	files, err := testutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("\nExpected a nil error but got %q.", err)
	}

	data := files.Get(filepath.Base(f.Name()))
	if len(data) == 0 {
		t.Fatal("\nExpected a not empty data.")
	}
}

func TestReadDirError(t *testing.T) {
	_, err := testutil.ReadDir("unknown")
	if err == nil {
		t.Error("\nExpected an error but got nil.")
	}
}
