/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"fmt"
	"github.com/chalet/cli/logger"
	"github.com/chalet/cli/utils"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the dev command for your project",
	Long: `Used to run the server locally on the chalet container. For example:
npm run dev`,
	Run: func(cmd *cobra.Command, args []string) {
		err := run()
		if err != nil {
			logger.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() error {
	config, err := utils.ReadConfig()
	if err != nil {
		return err
	}
	err = utils.CheckDockerContainerExists(config)
	if err != nil {
		return err
	}
	err = runCommand(config)
	if err != nil {
		return err
	}
	return nil
}

func runCommand(config *utils.Config) error {
	if config.ExposedPort == "" {
		config.ExposedPort = config.ServerPort
	}

	logger.Info(fmt.Sprintf("Running dev environment in %s...\n", config.ExposedPort))

	err := utils.Execute(config.Name, config.Commands.Run)
	if err != nil {
		return err
	}
	return nil
}
