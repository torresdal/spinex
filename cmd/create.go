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
  "crypto/tls"
  "gopkg.in/resty.v0"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
    // resty.SetDebug(true)

    cert1, err := tls.LoadX509KeyPair("/Users/jont/work/milesprojects/nocms/setup/spinnaker/client.pem", "/Users/jont/work/milesprojects/nocms/setup/spinnaker/private.pem")

    if err == nil {
      resty.SetCertificates(cert1)
      resp, err := resty.R().Get("https://deploy.milescloud.io:8084/applications")

      fmt.Println("\nError: %v", err)
      if err == nil {
        fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
        fmt.Printf("\nResponse Status: %v", resp.Status())
        fmt.Printf("\nResponse Time: %v", resp.Time())
        fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
        fmt.Printf("\nBody: %v", resp)

        var props []Property
        er := resty.Unmarshal(resp, &props)
        if er != nil {
            panic(er)
        } else {
            fmt.Println(props)
        }
      }
    } else {
      fmt.Println("\nError: %v", err)
    }
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
