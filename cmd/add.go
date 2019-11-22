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
	"fmt"
	"os"
	"io"
	"net/http"
	"path/filepath"

	"github.com/spf13/cobra"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Source URL]", // [Destination PATH]",
	Short: "Fetch goga-module and add it to the project.",
	Long: `Fetch goga-module from Source URL and put it as file into current directory.`,
  Args: cobra.RangeArgs(1,1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		filename := filepath.Base(url)
		fmt.Println("Fetch " + url + " into ./" + filename)

    if err := DownloadFile(filename, url); err != nil {
        panic(err)
    }

    // https://github.com/dapi/elements/blob/master/spinner.js
    // https://raw.githubusercontent.com/dapi/elements/master/spinner.js
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
