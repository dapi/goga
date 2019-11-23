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
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
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
	var re = regexp.MustCompile(`[^ ]+\s+goga\s+([^# ]+)$`)
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
		fmt.Println("push called", url)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
