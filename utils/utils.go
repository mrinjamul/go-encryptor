/*Package utils ...
 *
 */
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"syscall"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/ssh/terminal"
)

// AppName is the application name
var AppName = "go-encryptor"

// GetVersion return code of the application
func GetVersion() string {
	version := "1.0.0"
	return version
}

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

// ErrorLogger logs error
func ErrorLogger(err error) {
	log.Println(err)
}

// PromptTermPass takes password as user input
func PromptTermPass() ([]byte, error) {
	fmt.Print("password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		ErrorLogger(err)
		return []byte{}, err
	}
	return bytePassword, nil
}

// GetFileNameExt simplify filename for use (Note: only 3 char ext)
func GetFileNameExt(file string) (filename, extension string) {
	if len(file) > 4 && file[len(file)-4:len(file)-3] == "." {
		filename = file[0 : len(file)-4]
		extension = file[len(file)-3:]
	} else {
		filename = file
		extension = ""
	}
	return filename, extension
}

//ReadFile returns file data in bytes
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

//SaveFile save data to a file
func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
