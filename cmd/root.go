/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "chalet",
	Short: "Run dev environment in containers",
	Long: `Provides a streamlined way to set up and manage development
environments using Docker containers. 

With Chalet, you can run all your development tools, languages, and applications
in isolated containers, eliminating the need to install any languages or tools directly
on your local machine.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
