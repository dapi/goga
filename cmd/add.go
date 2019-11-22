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

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [Source URL] [Destination PATH]",
	Short: "Fetch goga-module and add it to the project",
	Long: `Fetch goga-module from Source URL and put it as file into Destination PATH. 
Use current directory if Destination PATH is not specified`,
  Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

var Region string

func init() {
	rootCmd.AddCommand(addCmd)

  // addCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
  // addCmd.MarkFlagRequired("region")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
  // addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
