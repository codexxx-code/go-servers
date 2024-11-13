package testutil

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var basepath struct {
	once  sync.Once
	value string
}

// Basepath returns the path of the root dir.
func Basepath() string {
	basepath.once.Do(func() {
		const skip int = 2

		_, file, _, _ := runtime.Caller(0)

		basepath.value = func(path string, n int) string {
			for i := 0; i < n; i++ {
				path = filepath.Dir(path)
			}

			return path
		}(file, skip)
	})

	return basepath.value
}

// Testdata returns the path of the testdata dir.
func Testdata() string {
	return filepath.Join(Basepath(), "testdata")
}

// Files is a cached files.
type Files struct {
	cached map[string][]byte
}

// newFiles allocate and returns a new Files.
func newFiles() *Files {
	return &Files{make(map[string][]byte)}
}

// Store stores the contents of a file by its name.
func (f Files) Store(name string, data []byte) {
	f.cached[name] = data
}

// Get returns the contents of the file by its name.
func (f Files) Get(name string) []byte {
	return f.cached[name]
}

// ReadDir returns cached files from the dir.
func ReadDir(dir string) (*Files, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := newFiles()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return nil, err
		}

		files.Store(name, data)
	}

	return files, nil
}
