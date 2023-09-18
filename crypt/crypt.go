package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/scrypt"
)

var (
	// MethodsAvailable is a list of available methods
	MethodsAvailable = []string{"aes", "aes256", "chacha20", "chacha20poly1305", "none"}
)

// XChaCha variables
var (
	SaltSize   = 32         // in bytes
	NonceSize  = 24         // in bytes. taken from aead.NonceSize()
	KeySize    = uint32(32) // KeySize is 32 bytes (256 bits).
	KeyTime    = uint32(5)
	KeyMemory  = uint32(1024 * 64) // KeyMemory in KiB. here, 64 MiB.
	KeyThreads = uint8(4)
	//chunkSize  = 1024 * 32 // chunkSize in bytes. here, 32 KiB.
)

// deriveKey hash the RAW password
func deriveKeyAES(password, salt []byte) ([]byte, []byte, error) {
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
	key, salt, err := deriveKeyAES(key, nil)
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
	key, _, err := deriveKeyAES(key, salt)
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
	key, _, err := deriveKeyAES(key, salt)
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

// ChaCha20Encrypt encrypts data with ChaCha20 Poly Algorithm and returns encrypted data and errors
func ChaCha20Encrypt(key, data []byte) ([]byte, error) {
	// Generate a new random 24-byte nonce.
	salt := make([]byte, SaltSize)
	if n, err := rand.Read(salt); err != nil || n != SaltSize {
		return nil, err
	}

	// Derive key from password
	key = argon2.IDKey(key, salt, KeyTime, KeyMemory, KeyThreads, KeySize)

	// Create a new ChaCha20 Poly1305 AEAD.
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	// Generate a new random 24-byte nonce.
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+aead.Overhead())

	if n, err := rand.Read(nonce); err != nil || n != NonceSize {
		return nil, err
	}

	// Encrypt data
	cipherdata := aead.Seal(nil, nonce, data, nil)

	// Append salt to data
	cipherdata = append(cipherdata, salt...)

	// Append nonce to cipherdata
	cipherdata = append(cipherdata, nonce...)

	return cipherdata, nil
}

// ChaCha20Decrypt decrypts data with ChaCha20 Poly Algorithm and returns decrypted data and errors
func ChaCha20Decrypt(key, data []byte) ([]byte, error) {
	// Extract nonce from data
	nonce := data[len(data)-NonceSize:]
	data = data[:len(data)-NonceSize]

	// Extract salt from data
	salt := data[len(data)-SaltSize:]
	data = data[:len(data)-SaltSize]

	// Derive key from password
	key = argon2.IDKey(key, salt, KeyTime, KeyMemory, KeyThreads, KeySize)

	// Create a new ChaCha20 Poly1305 AEAD.
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	// Decrypt data
	plaindata, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return plaindata, nil
}

// ChaCha20VerifyKey verify the password
func ChaCha20VerifyKey(key, data []byte) (bool, error) {
	// Extract nonce from data
	nonce := data[len(data)-NonceSize:]
	data = data[:len(data)-NonceSize]

	// Extract salt from data
	salt := data[len(data)-SaltSize:]
	data = data[:len(data)-SaltSize]

	// Derive key from password
	key = argon2.IDKey(key, salt, KeyTime, KeyMemory, KeyThreads, KeySize)

	// Create a new ChaCha20 Poly1305 AEAD.
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return false, err
	}

	// Decrypt data
	_, err = aead.Open(nil, nonce, data, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Encrypt encrypts data with given algorithm and returns encrypted data and errors
func Encrypt(data []byte, algorithm string, key []byte) ([]byte, error) {
	switch algorithm {
	case "aes":
		return AESEncrypt(key, data)
	case "chacha20":
		return ChaCha20Encrypt(key, data)
	default:
		return nil, errors.New("unknown algorithm")
	}
}

// Decrypt decrypts data with given algorithm and returns decrypted data and errors
func Decrypt(data []byte, algorithm string, key []byte) ([]byte, error) {
	switch algorithm {
	case "aes":
		return AESDecrypt(key, data)
	case "chacha20":
		return ChaCha20Decrypt(key, data)
	default:
		return nil, errors.New("unknown algorithm")
	}
}

// VerifyKey verify the password
func VerifyKey(data []byte, algorithm string, key []byte) (bool, error) {
	switch algorithm {
	case "aes":
		return AESVerifyKey(key, data)
	case "chacha20":
		return ChaCha20VerifyKey(key, data)
	default:
		return false, errors.New("unknown algorithm")
	}
}
