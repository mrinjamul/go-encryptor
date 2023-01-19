package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

// deriveKey hash the RAW password
func deriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, 1<<16, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}

// AESEncrypt returns encrypted data and errors
func AESEncrypt(key, data []byte) ([]byte, error) {
	key, salt, err := deriveKey(key, nil)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}
	cipherdata := gcm.Seal(nonce, nonce, data, nil)
	cipherdata = append(cipherdata, salt...)
	return cipherdata, nil
}

// AESDecrypt returns decrypted data and errors
func AESDecrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]
	key, _, err := deriveKey(key, salt)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, cipherdata := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaindata, err := gcm.Open(nil, nonce, cipherdata, nil)
	if err != nil {
		return nil, err
	}
	return plaindata, nil
}

// AESVerifyKey verify the password
func AESVerifyKey(key, data []byte) (bool, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]
	key, _, err := deriveKey(key, salt)
	if err != nil {
		return false, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return false, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return false, err
	}

	nonce, data := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	_, err = gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
