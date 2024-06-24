/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"bytes"
	"fmt"
	"github.com/chalet/cli/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the dev command for your project",
	Long: `Used to run the server locally on the chalet container. For example:
npm run dev`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
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

func run() {
	config, err := utils.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	exists := utils.CheckDockerContainerExists(config.Name)
	if !exists {
		fmt.Println("Container doesn't exist! Run chalet install to start the container")
		return
	}
	err = runCommand(config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runCommand(config *utils.Config) error {
	fmt.Println("Running dev environment...")

	cmdArgs := []string{"exec", config.Name, "sh", "-c", fmt.Sprintf("cd app && %s", config.Commands.Run)}
	cmd := exec.Command("docker", cmdArgs...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}
