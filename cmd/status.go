/*
Copyright © 2019 Danil Pismenny <danil@brandymint.ru>

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

	"github.com/SpectraLogic/go-gitignore"
	"github.com/spf13/cobra"
)

// Minimal possible file size
const MIN_FILE_SIZE = int64(len("/ goga http://ya.ru/1.js"))

func FilePathWalkDir(root string) ([]string, error) {
	gi, err := gitignore.NewGitIgnore("./.gitignore")
	CheckIfError(err)
	var files []string
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && (info.Name() == ".git" || (info.Name()[0] == '.' && len(info.Name()) > 1)) {
			return filepath.SkipDir
		}
		if !info.IsDir() && !gi.Match(path, false) {
			extension := filepath.Ext(path)
			if _, ok := Formats[extension]; ok {
				fi, err := os.Stat(path)
				CheckIfError(err)
				size := fi.Size()
				if size >= MIN_FILE_SIZE {
					firstLint := ReadFirstLine(path)
					if strings.Contains(firstLint, " goga ") {
						files = append(files, path)
					}
				}
			}
		}
		return nil
	})
	return files, err
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Scan every file in subdirectory and show its status",
	Long: `

> goga status`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Scanning.. ")
		files, err := FilePathWalkDir(".")
		CheckIfError(err)
		fmt.Println(len(files), "goga-files found")
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
