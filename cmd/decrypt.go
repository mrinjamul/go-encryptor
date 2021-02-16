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

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"de"},
	Short:   "",
	Long:    ``,
	Run:     decryptRun,
}

func decryptRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Too short argument")
		fmt.Println("Usage: " + utils.AppName + " encrypt [filename]")
		os.Exit(0)
	}
	encryptedfileName := args[0]

	filename, _ := utils.GetFileNameExt(encryptedfileName)

	password, err := utils.PromptTermPass()
	if err != nil {
		utils.ErrorLogger(err)
		os.Exit(1)
	}

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

	data, err := utils.AESDecrypt(password, encryptedData)
	if err != nil {
		fmt.Println("Error: Wrong Password")
		os.Exit(0)
	}

	encryptedExt, data := data[len(data)-3:], data[:len(data)-3]

	if string(encryptedExt) != "ger" {
		utils.SaveFile(filename+"."+string(encryptedExt), data)
	} else {
		utils.SaveFile(filename, data)
	}
	fmt.Println(filename + " decrypted successfully.")
}

var (
	keepdeOpt bool
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
}
