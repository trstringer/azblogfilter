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

	"github.com/trstringer/azblogfilter/internal/cache"
)

var (
	cfgFile          string
	keywordsFilter   string
	categoriesFilter string
	sinceFilter      string
	useCache         bool
	cacheLocation    string
	outputFormat     string
	getVersion       bool
)

var outputOptions = []string{"json", "csv"}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azblogfilter",
	Short: "Filter Azure update blog posts",
	Long: `Filter posts by keywords and categories.
You can also specify the time filter by using
a built-in caching mechanism or by specifying
the time manually

Examples:
	Get all blog posts in the past 7 days.

	$ azblogfilter --since -7d

	Get all blog posts since last run that have
	Kubernetes or Linux in the title.

	$ azblogfilter --cache --keywords "kubernetes,linux"

	Get all blog posts that have Linux in the title or
	have the DevOps category.

	$ azblogfilter --cache --keywords linux --categories devops`,
	Run: func(cmd *cobra.Command, args []string) {
		if getVersion {
			fmt.Println(version)
			os.Exit(0)
		}
		err := validateArgs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		since, err := effectiveSinceTime()
		if err != nil {
			fmt.Printf("Error getting last time: %v\n", err)
			os.Exit(1)
		}

		posts, err := getBlogPosts(since, keywordsFilter, categoriesFilter)
		if err != nil {
			fmt.Printf("Error getting blog posts: %v\n", err)
			os.Exit(1)
		}

		output, err := formatBlogPosts(outputFormat, posts)
		if err != nil {
			fmt.Printf("Error formatting output: %v\n", err)
			os.Exit(1)
		}

		if useCache {
			cachePath, err := realCachePath()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = cache.SetLastCachedTime(cachePath, time.Now().UTC())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Print(output)
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
	rootCmd.Flags().StringVarP(&sinceFilter, "since", "s", "", "filter post with systemd time (man 7 systemd.time)")
	rootCmd.Flags().BoolVar(&useCache, "cache", false, "use cached value")
	rootCmd.Flags().StringVar(&cacheLocation, "cache-location", "~/.azblogfilter", "location for cache")
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "output format, json or csv")
	rootCmd.Flags().BoolVarP(&getVersion, "version", "v", false, "get version")
}

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
