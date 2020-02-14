package test

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// ReadFileAsString is a testing helper, that returns the file if possible. If
// there is an error, it will call t.Fatal(), ending the test.
//
// Example:
//     content := test.ReadFileAsString(t, "folderA", "folderB", "file.txt")
func ReadFileAsString(t *testing.T, path ...string) string {
	t.Helper()
	b := ReadFileAsBytes(t, path...)
	s := string(b)
	return s
}

// ReadFileAsBytes is a testing helper, that returns the file if possible. If
// there is an error, it will call t.Fatal(), ending the test.
func ReadFileAsBytes(t *testing.T, path ...string) []byte {
	t.Helper()
	joinedPath := filepath.Join(path...) // relative path
	bytes, err := ioutil.ReadFile(joinedPath)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
