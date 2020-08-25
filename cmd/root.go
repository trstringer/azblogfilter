// Package cmd is the base root command.
/*
Copyright © 2020 Thomas Stringer

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
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/trstringer/go-systemd-time/systemdtime"
)

var (
	cfgFile          string
	keywordsFilter   string
	categoriesFilter string
	sinceFilter      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azblogfilter",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		adjustedTime, err := systemdtime.AdjustTime(&now, sinceFilter)
		if err != nil {
			fmt.Printf("Error converting systemd time: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Adjusted time is %v\n", adjustedTime)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azblogfilter.yaml)")

	rootCmd.Flags().StringVarP(&keywordsFilter, "keywords", "k", "", "keywords filter (case insensitive)")
	rootCmd.Flags().StringVarP(&categoriesFilter, "categories", "c", "", "categories filter (case insensitive)")
	rootCmd.Flags().StringVarP(&sinceFilter, "since", "s", "-7d", "filter post with systemd time (man 7 systemd.time). Default -7d")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".azblogfilter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".azblogfilter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
