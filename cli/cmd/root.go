// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "spinex",
	Short: "Spinex is a simple CLI for Spinnaker.",
	Long: `Spinex is a simple CLI for Spinnaker.

To get autocompletion, run 'sudo spinex gen autocomplete'`,
	Run: func(cmd *cobra.Command, args []string) {
    if showVersion {
      fmt.Println("Spinex version 0.1")
    } else {
      cmd.Help()
    }
  },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var (
  showVersion bool
  account string
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spinex.yaml)")
  RootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show current version of Spinex")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".spinex") // name of config file (without extension)
  viper.AddConfigPath("/etc/spinex/")
	viper.AddConfigPath("$HOME/")  // adding home directory as first search path
  viper.AddConfigPath(".")
  viper.SetConfigType("yaml")
	// viper.AutomaticEnv()          // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
    fmt.Println()
	} else {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
}
