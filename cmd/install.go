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
	"strings"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all needed dependencies for your project",
	Long: `Installs all the necessary dependencies, based on the command
configured on the chalet.yaml file.For example:
chalet install`,
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func install() {
	config, err := utils.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	exists := checkDockerContainerExists(config.Name)
	if !exists {
		err := createContainer(config)
		if err != nil {
			fmt.Println(err)
		}
		err = startContainer(config)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = installDependencies(config)
	if err != nil {
		fmt.Println(err)
	}
}

func checkDockerContainerExists(containerName string) bool {
	cmd := exec.Command("docker", "ps", "-a", "--filter", fmt.Sprintf("name=^%s$", containerName), "--format", "{{.Names}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running Docker command:", err)
		return false
	}

	return bytes.Contains(out.Bytes(), []byte(containerName))
}

func createContainer(config *utils.Config) error {
	cmd := exec.Command("docker", "create", "-v", "$(pwd):/app", "-p", fmt.Sprintf("7300:%d", config.ServerPort), "--name", config.Name, fmt.Sprintf("%s:%s", config.Lang, config.Version))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func startContainer(config *utils.Config) error {
	cmd := exec.Command("docker", "start", config.Name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func installDependencies(config *utils.Config) error {
	command := fmt.Sprintf("exec -it %s cd app & %s", config.Name, config.Commands.Install)
	cmd := exec.Command("docker", strings.Split(command, " ")...)
	err := cmd.Run()
	fmt.Println(command)
	if err != nil {
		return err
	}
	return nil
}
