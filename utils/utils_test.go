package utils

import (
	"testing"
)

// TestGetVersion tests
func TestGetVersion(t *testing.T) {
	out := GetVersion()
	if out == "" || len(out) == 0 {
		t.Errorf("Want strings but got nil")
	}
}

// TestGetFileNameExt test for GetFileNameExt
func TestGetFileNameExt(t *testing.T) {
	testcases := []struct {
		name         string
		fullFileName string
		filename     string
		extension    string
	}{
		{"File with Extension", "test.jpg", "test", "jpg"},
		{"File without Extension", "test", "test", ""},
		{"File with 2 latter Extension", "test.md", "test", "md"},
		{"File name smaller than 4", "xyz", "xyz", ""},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fileName, ext := GetFileNameExt(testcase.fullFileName)
			if fileName != testcase.filename || ext != testcase.extension {
				t.Errorf("%v should be '%v' with '%v'; but got '%v' with '%v'",
					testcase.fullFileName, testcase.filename,
					testcase.extension, fileName, ext)
			}
		})
	}
}
