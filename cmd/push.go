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
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func ReadFirstLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var firstLine string
	firstLine, err = reader.ReadString('\n')

	return firstLine
}

func FetchUrlFromComment(comment string) string {
	re := regexp.MustCompile(`[^ ]+\s+goga\s+([^# ]+)$`)
	return re.ReplaceAllString(comment, `$1`)
}

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push <file>",
	Short: "Push changes of goga-modules into source repository",
	Long: `Push specified file into its repository. For example:

# Push single file
> goga push ./app/javascripts/spinner.js
`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		firstLint := ReadFirstLine(file)
		url := FetchUrlFromComment(firstLint)
		PushFileToRemoteRepository(file, url)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func CopyRemovingMagicComment(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	line := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if line > 0 {
			out.WriteString(scanner.Text() + "\n")
		}
		line += 1
	}
	return out.Close()
}

func GetSubdirectoryFromUrl(url string) string {
	re := regexp.MustCompile(`^https://github.com/[^\/]+/[^\/]+/blob/([^\/]+)/(.+)\n?$`)
	// $1 - branch
	return re.ReplaceAllString(url, `$2`)
}

func GetRepoFromUrl(url string) string {
	re := regexp.MustCompile(`^https://github.com/([^\/]+/[^\/]+)/blob/([^\/]+)/(.+)\n?$`)
	return re.ReplaceAllString(url, `git@github.com:$1.git`)
}

func PushFileToRemoteRepository(file string, url string) error {
	tempDir, err := ioutil.TempDir("", "goga")
	CheckIfError(err)
	defer os.RemoveAll(tempDir)

	destination_file := GetSubdirectoryFromUrl(url)
	fmt.Print("Checking ", file)

	// TODO Get repo from url
	var repo = GetRepoFromUrl(url)

	fmt.Print(" to ", repo+"/"+destination_file)

	r, err := git.PlainClone(tempDir, false, &git.CloneOptions{URL: repo})
	CheckIfError(err)

	fmt.Print(" ")
	destination_file_path := tempDir + "/" + destination_file

	// TODO Remove magic-comment
	CopyRemovingMagicComment(file, destination_file_path)

	w, err := r.Worktree()
	CheckIfError(err)

	status, err := w.Status()
	CheckIfError(err)

	if len(status) == 0 {
		fmt.Println("nothing changed, skip")
	} else {
		CommintAndPush(r, w, tempDir, destination_file, destination_file_path)
	}

	return err
}

func CommintAndPush(r *git.Repository, w *git.Worktree, tempDir string, destination_file string, destination_file_path string) {
	username, err := gitconfig.Username()
	email, err := gitconfig.Email()
	commitMessage := fmt.Sprintf("Update %s by goga", destination_file)

	_, err = w.Add(destination_file)
	CheckIfError(err)

	commit, err := w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})
	CheckIfError(err)

	fmt.Print("commit, ")
	_, err = r.CommitObject(commit)
	CheckIfError(err)

	fmt.Print("pushing.. ")

	err = r.Push(&git.PushOptions{})
	CheckIfError(err)
	fmt.Println("done.")
	return
}
