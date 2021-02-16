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

// TestderiveKey hash the RAW password
func TestDeriveKey(t *testing.T) {
	password := []byte("password")
	key, salt, err := deriveKey(password, nil)
	if len(key) == 0 || len(salt) == 0 {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestAESEncrypt tests
func TestAESEncrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	cipherdata, err := AESEncrypt(password, data)
	if cipherdata == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestAESDecrypt tests
func TestAESDecrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	cipherdata, _ := AESEncrypt(password, data)
	plaindata, err := AESDecrypt(password, cipherdata)
	if plaindata == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestVerifyKey verify the password
func TestAESVerifyKey(t *testing.T) {
	password := []byte("password")
	fakepassword := []byte("fakepassword")
	data := []byte("Data")
	cipherdata, _ := AESEncrypt(password, data)
	res, err := AESVerifyKey(password, cipherdata)
	if err != nil {
		t.Errorf("Want erors to be nil but got %v\n", err)
	}
	if res != true {
		t.Errorf("Want true but got false")
	}
	res, err = AESVerifyKey(fakepassword, cipherdata)
	if err == nil {
		t.Errorf("Want erors but got nil")
	}
	if res != false {
		t.Errorf("Want false but got true")
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
