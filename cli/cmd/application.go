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

// applicationsCmd represents the applications command
var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Interacts with Spinnaker applications",
	Long: "",
}

var appListCmd = &cobra.Command {
  Use: "list",
  Short: "Lists applications",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])
    client.Applications(cl)
  },
}

var appCreateCmd = &cobra.Command {
  Use: "create NAME EMAIL ACCOUNTS",
  Short: "Create application",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			cmd.Help()
			return
		}

		var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])
    client.CreateApplication(cl, args[0], args[1], args[2], appCreateCloudProviders, appCreatePort, appCreateDescription)
  },
}

var appDeleteCmd = &cobra.Command {
  Use: "delete NAME",
  Short: "Delete application",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
      cmd.Help()
      return
    }

    var spinnaker = viper.GetStringMapString("spinnaker")

    cl := client.NewClient(spinnaker["host"], spinnaker["x509certfile"], spinnaker["x509keyfile"])
    client.DeleteApplication(cl, args[0])
  },
}

var (
  appCreateDescription string
  appCreatePort string
  appCreateCloudProviders string
)

func init() {
	RootCmd.AddCommand(applicationCmd)

  applicationCmd.AddCommand(appListCmd)
  applicationCmd.AddCommand(appCreateCmd)
  applicationCmd.AddCommand(appDeleteCmd)

  appCreateCmd.Flags().StringVarP(&appCreateDescription, "descr", "d", "", "Description of application")
  appCreateCmd.Flags().StringVarP(&appCreatePort, "port", "p", "", "Instance port")
  appCreateCmd.Flags().StringVar(&appCreateCloudProviders, "providers", "", "Cloud Providers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
