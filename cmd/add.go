/*
Copyright © 2019 Danil Pismenny <danil@brandymint.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

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

func GenerateMagicComment(url string, ext string) string {
	format := Formats[ext]
	if len(format) == 0 {
		panic(fmt.Sprintf("I don't known how to add comments for this extension - '%s'", ext))
	}
	return fmt.Sprintf(format, "goga", url)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filename string, url string, original_url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	ext := filepath.Ext(filename)
	line := 0
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if line == 0 {
			out.WriteString(GenerateMagicComment(original_url, ext))
		}
		out.WriteString("\n" + scanner.Text())
		line += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Write the body to file
	//_, err = io.Copy(out, resp.Body)
	return err
}

func replaceGithubDirectLink(url string) string {
	var re = regexp.MustCompile(`^(https://github.com/)([^/]+/[^/]+/)blob/(.+)$`)
	s := re.ReplaceAllString(url, `https://raw.githubusercontent.com/$2$3`)
	return s
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Source URL]", // [Destination PATH]",
	Short: "Fetch goga-module and add it to the project.",
	Long:  `Fetch goga-module from Source URL and put it as file into current directory.`,
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		original_url := args[0]
		url := replaceGithubDirectLink(original_url)
		filename := filepath.Base(url)
		fmt.Println("Fetch " + url + " into ./" + filename)

		if err := DownloadFile(filename, url, original_url); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
