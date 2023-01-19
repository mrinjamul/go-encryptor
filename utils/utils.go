/*Package utils ...
 *
 */
package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// AppName is the application name
var AppName = "go-encryptor"

// ErrorLogger logs error
func ErrorLogger(err error) {
	log.Println(err)
}

// PromptTermPass takes password as user input
func PromptTermPass(promptText string) ([]byte, error) {
	fmt.Print(promptText)
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
	} else if len(file) > 3 && file[len(file)-3:len(file)-2] == "." {
		filename = file[0 : len(file)-3]
		extension = file[len(file)-2:]
	} else {
		filename = file
		extension = ""
	}
	return filename, extension
}

// ReadFile returns file data in bytes
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// SaveFile save data to a file
func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// IsDir is to check if its a dir
func IsDir(path string) (bool, error) {
	out, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return out.IsDir(), nil
}

// ConfirmPrompt will prompt to user for yes or no
func ConfirmPrompt(message string) bool {
	var response string
	fmt.Print(message + " (yes/no) :")
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		return false
	}
}
