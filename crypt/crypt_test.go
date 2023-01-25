package crypt

import (
	"testing"
)

// TestderiveKey hash the RAW password
func TestDeriveKey(t *testing.T) {
	password := []byte("password")
	key, salt, err := deriveKeyAES(password, nil)
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

// TestChaCha20Encrypt tests
func TestChaCha20Encrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	cipherdata, err := ChaCha20Encrypt(password, data)
	if cipherdata == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestChaCha20Decrypt tests
func TestChaCha20Decrypt(t *testing.T) {
	password := []byte("password")
	data := []byte("Data")
	cipherdata, _ := ChaCha20Encrypt(password, data)
	plaindata, err := ChaCha20Decrypt(password, cipherdata)
	if plaindata == nil {
		t.Errorf("Want strings but got nil")
	}
	if err != nil {
		t.Errorf("Want nil but got errors")
	}
}

// TestChaCha20VerifyKey verify the password
func TestChaCha20VerifyKey(t *testing.T) {
	password := []byte("password")
	fakepassword := []byte("fakepassword")
	data := []byte("Data")
	cipherdata, _ := ChaCha20Encrypt(password, data)
	res, err := ChaCha20VerifyKey(password, cipherdata)
	if err != nil {
		t.Errorf("Want erors to be nil but got %v\n", err)
	}
	if res != true {
		t.Errorf("Want true but got false")
	}
	res, err = ChaCha20VerifyKey(fakepassword, cipherdata)
	if err == nil {
		t.Errorf("Want erors but got nil")
	}
	if res != false {
		t.Errorf("Want false but got true")
	}
}
