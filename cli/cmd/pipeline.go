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
  "fmt"
	"github.com/spf13/cobra"
  "github.com/torresdal/spinex/client"
  "github.com/spf13/viper"
)

var (
  pipeStartTag   string
)

// pipelineCmd represents the pipeline command
var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Interact with pipelines for an application",
	Long: "",
}

var pipeListCmd = &cobra.Command{
	Use:   "list APP_NAME",
	Short: "List pipelines",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
      cmd.Help()
      return
    }

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])
    client.Pipelines(cl, args[0])
	},
}

var pipeCreateCmd = &cobra.Command{
	Use:   "create APP_NAME",
	Short: "Create pipeline",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("asdfasdf")
	},
}

var pipeDeleteCmd = &cobra.Command{
	Use:   "delete APP_NAME NAME",
	Short: "Delete pipeline",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("asdfasdf")
	},
}

var pipeStartCmd = &cobra.Command{
	Use:   "start APP_NAME PIPELINE",
	Short: "Start pipeline execution. Currently only the Docker Registry Trigger is supported.",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
      cmd.Help()
      return
    }

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])
    client.StartPipeline(cl, args[0], args[1], pipeStartTag)
	},
}

func init() {
	RootCmd.AddCommand(pipelineCmd)
  pipelineCmd.AddCommand(pipeListCmd)
  pipelineCmd.AddCommand(pipeCreateCmd)
  pipelineCmd.AddCommand(pipeDeleteCmd)
  pipelineCmd.AddCommand(pipeStartCmd)

  pipeStartCmd.Flags().StringVarP(&pipeStartTag, "tag", "t", "", "Which Docker image tag to use for pipeline. If omitted, a prompt of max 20 tags will let you choose.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
