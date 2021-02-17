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

	"github.com/mrinjamul/go-encryptor/utils"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Aliases: []string{"en", "e"},
	Short:   "Encrypt file using specified method",
	Long:    `Encrypt file using specified method. (Default: AES)`,
	Run:     encryptRun,
}

func encryptRun(cmd *cobra.Command, args []string) {
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

	fileName := args[0]
	filename, extension := utils.GetFileNameExt(fileName)
	if extension == "" {
		extension = "ger"
	}
	encryptFileName := filename + ".aes"

	password, err := utils.PromptTermPass("Password: ")
	if err != nil {
		utils.ErrorLogger(err)
		os.Exit(1)
	}

	if len(password) < 5 {
		fmt.Println("Warning: Insecure password")
		fmt.Println("You should use password with more than 5 characters")
	}

	verifyPassword, err := utils.PromptTermPass("Verify Password: ")
	if err != nil {
		utils.ErrorLogger(err)
		os.Exit(1)
	}

	if string(verifyPassword) != string(password) {
		fmt.Println("Error: Both password is not same")
		os.Exit(1)
	}

	data, err := utils.ReadFile(fileName)
	if err != nil {
		utils.ErrorLogger(err)
		os.Exit(1)
	}
	data = append(data, extension...)

	encryptdata, err := utils.AESEncrypt(password, data)
	if err != nil {
		utils.ErrorLogger(err)
	}

	utils.SaveFile(encryptFileName, encryptdata)
	fmt.Println(fileName + " encrypted successfully.")

	if !keepenOpt {
		err := os.Remove(args[0])
		if err != nil {
			utils.ErrorLogger(err)
		}
	}
}

var (
	keepenOpt bool
)

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encryptCmd.Flags().BoolVarP(&keepenOpt, "keep", "k", false, "Keep uncrypted file")
}
