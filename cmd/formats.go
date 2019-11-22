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

	"github.com/spf13/cobra"
)

var CommentPrefix = "goga"
var Formats = map[string]string{
	".c":     "// %s %s",
	".js":    "// %s %s",
	".go":    "// %s %s",
	".java":  "// %s %s",
	".php":   "// %s %s",
	".slim":  "// %s %s",
	".haml":  "// %s %s",
	".rb":    "# %s %s",
	".py":    "# %s %s",
	".pl":    "# %s %s",
	".sql":   "-- %s %s",
	".swift": "-- %s %s",
	".xml":   "<!-- %s %s -->",
	".html":  "<!-- %s %s -->",
}

// formatsCmd represents the formats command
var formatsCmd = &cobra.Command{
	Use:   "formats",
	Short: "Pring list of available file extensions and comments syntax",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List of available file extensions and its comments syntax:\n")
		for extension, format := range Formats {
			fmt.Printf(extension+"\t"+format+"\n", CommentPrefix, "URI")
		}
	},
}

func init() {
	rootCmd.AddCommand(formatsCmd)
}
