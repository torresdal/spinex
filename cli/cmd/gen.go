// Copyright © 2016 Jon Arild Torresdal <jon@torresdal.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "log"
  "fmt"
	"github.com/spf13/cobra"
)

var autocompleteTarget string
var autocompleteType string

// applicationsCmd represents the applications command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate files",
	Long: "",
}

var genAutocompleteCmd = &cobra.Command {
  Use: "autocomplete",
  Short: "Generate file for bash autocompletion",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    err := cmd.Root().GenBashCompletionFile(autocompleteTarget)
    if err != nil {
			log.Fatal(err)
		}

    fmt.Println("Bash completion file for Hugo saved to", autocompleteTarget)
  },
}

func init() {
	RootCmd.AddCommand(genCmd)
  genCmd.AddCommand(genAutocompleteCmd)

  genAutocompleteCmd.PersistentFlags().StringVarP(&autocompleteTarget, "completionfile", "", "/etc/bash_completion.d/spinex.sh", "Autocompletion file")
  genAutocompleteCmd.PersistentFlags().StringVarP(&autocompleteType, "type", "", "bash", "Autocompletion type (currently only bash supported)")
}
