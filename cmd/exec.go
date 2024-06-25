/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"bufio"
	"fmt"
	"github.com/chalet/cli/utils"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute your custom commands",
	Long: `The exec command allows you to run custom commands defined in your configuration file.
These commands are designed to simplify and automate various tasks within your project.

Examples:

1. Running a custom command defined in your chalet.yaml configuration file:
   $ chalet exec my_custom_command

2. Executing an arbitrary shell command within the Chalet container:
   $ chalet exec "echo Hello, World!"

The command first checks if the provided command exists in the custom commands defined in your
configuration file (chalet.yaml). If it exists, it executes the corresponding command.
If not, it treats the input as a regular shell command and executes it within the container.
`,
	Run: func(cmd *cobra.Command, args []string) {
		execHandler(args)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execHandler(args []string) {
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

	err = execCommand(config, strings.Join(args, " "))
	if err != nil {
		fmt.Println(err)
	}

	err = utils.StopContainer(config)
	if err != nil {
		fmt.Println(err)
	}
}

// Use the follwoing command to test log streaming"
// ./chalet exec 'bash -c "for i in {1..11}; do echo hola; sleep 0.2; done"'
func execCommand(config *utils.Config, args string) error {
	var commandToRun string
	if command, exists := config.CustomCommands[args]; exists {
		fmt.Println("Executing", command, "from chalet.yaml...")
		commandToRun = command
	} else {
		fmt.Println("Executing", args, "...")
		commandToRun = args
	}
	cmdArgs := []string{"exec", fmt.Sprintf("chalet-%s", config.Name), "sh", "-c", fmt.Sprintf("cd /chalet && %s", commandToRun)}
	cmd := exec.Command("docker", cmdArgs...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating stderr pipe: %v", err)
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start docker command: %v", err)
	}

	// Use a scanner to read and print stdout
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Use a scanner to read and print stderr
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Wait for command to finish
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("command finished with error: %v", err)
	}

	return nil
}
