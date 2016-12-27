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
	"github.com/spf13/cobra"
  "github.com/torresdal/spinex/client"
  "github.com/spf13/viper"
)

var execName string
var execStatuses string
var execSortBy string
var execDesc bool
var execLimit int
var execCancelReason string

// applicationsCmd represents the applications command
var executionCmd = &cobra.Command{
	Use:   "execution",
	Short: "Interacts with Spinnaker pipeline executions",
	Long: "",
}

var execListCmd = &cobra.Command {
  Use: "list APP",
  Short: "List pipeline exeuctions for application",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
			cmd.Help()
			return
		}

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])

    client.Executions(cl, args[0], execLimit, execStatuses, execSortBy, execDesc, execName)
  },
}

var execCancelCmd = &cobra.Command {
  Use: "cancel ID",
  Short: "Cancel pipeline exeuction",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
			cmd.Help()
			return
		}

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])

    client.CancelExecution(cl, args[0], execCancelReason)
  },
}

var execDeleteCmd = &cobra.Command {
  Use: "delete ID",
  Short: "Delete pipeline exeuction",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
			cmd.Help()
			return
		}

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])

    client.DeleteExecution(cl, args[0])
  },
}

var execInfoCmd = &cobra.Command {
  Use: "info ID",
  Short: "Get detailed information of pipeline exeuction",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
			cmd.Help()
			return
		}

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])

    client.ExecutionInfo(cl, args[0])
  },
}

func init() {
	RootCmd.AddCommand(executionCmd)

  executionCmd.AddCommand(execListCmd)
  executionCmd.AddCommand(execCancelCmd)
  executionCmd.AddCommand(execDeleteCmd)
  executionCmd.AddCommand(execInfoCmd)

  execListCmd.Flags().IntVarP(&execLimit, "limit", "l", 0, "Limit to n results per pipeline")
  execListCmd.Flags().StringVarP(&execStatuses, "statuses", "s", "", "Filter on statuses (SUCCEEDED,RUNNING,TERMINAL,CANCELED)")
  execListCmd.Flags().StringVar(&execSortBy, "sort", "START", "Sort by NAME, START, END or STATUS")
  execListCmd.Flags().BoolVar(&execDesc, "desc", false, "Used with --sort to sort ascending (default is descending)")
  execListCmd.Flags().StringVar(&execName, "name", "", "Only show pipelines with this name")

  execCancelCmd.Flags().StringVar(&execCancelReason, "reason", "", "Reason for cancellation")
}
