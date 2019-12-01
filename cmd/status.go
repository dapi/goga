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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/SpectraLogic/go-gitignore"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
)

// Minimal possible file size
const MIN_FILE_SIZE = int64(len("/ goga http://ya.ru/1.js"))

func FilePathWalkDir(root string) error {
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
						CheckFileStatus(path)

						files = append(files, path)
					}
				}
			}
		}
		return nil
	})
	return err
}

func CheckFileStatus(file string) {
	firstLint := ReadFirstLine(file)
	url := FetchUrlFromComment(firstLint)

	tempDir, err := ioutil.TempDir("", "goga")
	CheckIfError(err)
	defer os.RemoveAll(tempDir)

	destination_file := GetSubdirectoryFromUrl(url)
	fmt.Print("Found ", file, " checking")

	var repo = GetRepoFromUrl(url)

	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{URL: repo})
	CheckIfError(err)

	fmt.Print(" ")
	destination_file_path := tempDir + "/" + destination_file

	tmpfile_local, err := ioutil.TempFile("", "goga.local")
	//fmt.Println(tmpfile_local.Name())
	defer os.Remove(tmpfile_local.Name())

	tmpfile_remote, err := ioutil.TempFile("", "goga.remote")
	//fmt.Println(tmpfile_remote.Name())
	defer os.Remove(tmpfile_remote.Name())

	CopyWithoutMagicComment(file, tmpfile_local.Name())
	CopyWithoutMagicComment(destination_file_path, tmpfile_remote.Name())

	content_local, err := ioutil.ReadFile(tmpfile_local.Name())
	CheckIfError(err)

	content_remote, err := ioutil.ReadFile(tmpfile_remote.Name())
	CheckIfError(err)

	text_local := string(content_local)
	text_remote := string(content_remote)

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text_local, text_remote, false)

	if len(diffs) > 1 {
		fmt.Println("-", len(diffs)-1, "diffs found")
		fmt.Println(dmp.DiffPrettyText(diffs))
	} else {
		fmt.Println("- no changes")
	}

}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status <dir to scan>",
	Short: "Scan every file in subdirectories of specified directory and show its status",
	Long: `Scan every file in subdirectories of specified directory and show its status. 
Uses current directory if no arguments specified. For example:

> goga status`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		if len(args) == 0 {
			dir = "."
		} else {
			dir = args[0]
		}
		fmt.Println("Scanning directory:", dir)
		err := FilePathWalkDir(dir)
		CheckIfError(err)
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
