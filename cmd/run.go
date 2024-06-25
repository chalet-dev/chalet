/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"bytes"
	"fmt"
	"github.com/chalet/cli/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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
		return
	}
	err = utils.CheckDockerContainerExists(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = runCommand(config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func runCommand(config *utils.Config) error {
	if config.ExposedPort == "" {
		config.ExposedPort = "7300"
	}
	fmt.Println(fmt.Sprintf("Running dev environment in %s...", config.ExposedPort))

	// Create a channel to listen for termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	cmdArgs := []string{"exec", fmt.Sprintf("chalet-%s", config.Name), "sh", "-c", fmt.Sprintf("cd /chalet && %s", config.Commands.Run)}
	cmd := exec.Command("docker", cmdArgs...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	done := make(chan error, 1)

	go func() {
		done <- cmd.Run()
	}()

	select {
	case err := <-done:
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return err
		}
		return nil
	case sig := <-sigChan:
		fmt.Println("Stopping the container...")

		// Run docker stop
		stopCmd := exec.Command("docker", "stop", fmt.Sprintf("chalet-%s", config.Name))
		stopErr := stopCmd.Run()
		if stopErr != nil {
			fmt.Println("Failed to stop container:", stopErr)
			return stopErr
		}
		fmt.Println("Chalet stopped successfully.")
		return fmt.Errorf("process interrupted by signal: %v", sig)
	}
}
