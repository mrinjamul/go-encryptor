/*
Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mrinjamul/go-encryptor/crypt"
	"github.com/mrinjamul/go-encryptor/utils"
	twarper "github.com/mrinjamul/go-tar/tarwarper"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"de"},
	Short:   "Decrypt encrypted file using specified method",
	Long:    `Decrypt encrypted file using specified method. (Default: AES)`,
	Run:     decryptRun,
}

func decryptRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Too short argument")
		fmt.Println("Usage: " + utils.AppName + " encrypt [filename]")
		os.Exit(0)
	}

	if len(args) > 1 {
		fmt.Println("Error: Too many argument")
		fmt.Println("Usage: " + utils.AppName + " encrypt [filename]")
		os.Exit(0)
	}

	encryptedfileName := args[0]
	var filename string // filename without extension

	// check if file has no extension with filepath package
	exten := filepath.Ext(encryptedfileName)
	if exten == ".aes" {
		filename = strings.TrimSuffix(encryptedfileName, exten)
	} else {
		filename = encryptedfileName
		encryptedfileName = encryptedfileName + ".aes"
	}

	// check if file exists
	if _, err := os.Stat(encryptedfileName); os.IsNotExist(err) {
		fmt.Println("Error: No such file or directory")
		os.Exit(1)
	}

	var password []byte

	if passwordOpt == "" {
		p, err := utils.PromptTermPass("Password: ")
		password = p
		if err != nil {
			utils.ErrorLogger(err)
			os.Exit(1)
		}
	} else {
		password = []byte(passwordOpt)
	}

	// read encrypted data
	encryptedData, err := utils.ReadFile(encryptedfileName)
	if err != nil {
		utils.ErrorLogger(err)
		os.Exit(1)
	}

	// verifyOpt, _ := utils.AESVerifyKey(password, encryptedData)
	// if verifyOpt != true {
	// 	fmt.Println("Error: Wrong Password")
	// 	os.Exit(0)
	// }

	var data []byte
	// check if encryption method
	if methodOpt == "aes" {
		data, err = crypt.AESDecrypt(password, encryptedData)
		if err != nil {
			fmt.Println("Error: Wrong Password")
			os.Exit(0)
		}
	} else if methodOpt == "xchacha20" || methodOpt == "chacha20" {
		data, err = crypt.ChaCha20Decrypt(password, encryptedData)
		if err != nil {
			fmt.Println("Error: Wrong Password")
			os.Exit(0)
		}
	}

	encryptedExt, data := data[len(data)-3:], data[:len(data)-3]

	// print to stdout and return
	if stdoutOpt {
		fmt.Print(data)
		return
	}

	if strings.Contains(string(encryptedExt), "2") {
		encryptedExt = encryptedExt[:2]
	}

	if string(encryptedExt) != "ger" {
		if !strings.HasSuffix(filename, ".") {
			filename = filename + "."
		}
		utils.SaveFile(filename+string(encryptedExt), data)
	} else {
		utils.SaveFile(filename, data)
	}
	if string(encryptedExt) == "tez" {
		path, err := os.Getwd()
		if err != nil {
			utils.ErrorLogger(err)
		}
		err = twarper.ExtarctTar(path, filename+".tez")
		_ = os.Remove(filename + ".tez")
		if err != nil {
			utils.ErrorLogger(err)
			os.Exit(1)
		}
	}
	fmt.Println(filename + " decrypted successfully.")

	if !keepdeOpt {
		err := os.Remove(encryptedfileName)
		if err != nil {
			utils.ErrorLogger(err)
		}
	}
}

var (
	keepdeOpt bool
	stdoutOpt bool
)

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decryptCmd.Flags().BoolVarP(&keepdeOpt, "keep", "k", false, "Keep encrypted file")
	decryptCmd.Flags().BoolVarP(&stdoutOpt, "out", "o", false, "Print to stdout")
	decryptCmd.Flags().StringVarP(&passwordOpt, "password", "p", "", "Get password")
	// methods for encryption flag
	decryptCmd.Flags().StringVarP(&methodOpt, "method", "m", "aes", "Encryption method")
}
