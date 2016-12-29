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
  "os"
  "path/filepath"
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
  Short: "Generate file for bash autocompletion.",
  Long: "Generate file for bash autocompletion. Probably needs sudo: sudo spinex gen autocomplete",
  Run: func(cmd *cobra.Command, args []string) {
    err := os.MkdirAll(autocompleteTarget, os.ModePerm)
    if err != nil {
			log.Fatal(err)
		}

    path := filepath.Join(autocompleteTarget, "spinex.sh")
    err = cmd.Root().GenBashCompletionFile(path)
    if err != nil {
			log.Fatal(err)
		}

    fmt.Println("Bash completion file for Hugo saved to", path)
    fmt.Println("")
    fmt.Println("Restart terminal or run ")
    fmt.Println("")
    fmt.Printf("  . %s\n", path)
    fmt.Println("")
    fmt.Println("to enable autocompletion now.")
  },
}

var genSpinexConfig = &cobra.Command {
  Use: "config",
  Short: "Generate spinex configuration file",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Not implemented")
  },
}

func init() {
	RootCmd.AddCommand(genCmd)
  genCmd.AddCommand(genAutocompleteCmd)
  genCmd.AddCommand(genSpinexConfig)

  genAutocompleteCmd.PersistentFlags().StringVarP(&autocompleteTarget, "completiondir", "", "/etc/bash_completion.d", "Autocompletion target dir")
}
