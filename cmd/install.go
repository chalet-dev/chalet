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
		return
	}
	exists := utils.CheckDockerContainerExists(config.Name)
	if !exists {
		err := createContainer(config)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = startContainer(config)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = installDependencies(config)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createContainer(config *utils.Config) error {
	fmt.Println("Creating container...")
	cwd, err := os.Getwd()
	cmd := exec.Command("docker", "create", "-it", "-v", fmt.Sprintf("%s:/app", cwd), "-p", fmt.Sprintf("7300:%d", config.ServerPort), "--name", config.Name, fmt.Sprintf("%s:%s", config.Lang, config.Version))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func startContainer(config *utils.Config) error {
	fmt.Println("Starting container...")
	cmd := exec.Command("docker", "start", config.Name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func installDependencies(config *utils.Config) error {
	fmt.Println("Installing dependencies...")

	cmdArgs := []string{"exec", config.Name, "sh", "-c", fmt.Sprintf("cd app && %s", config.Commands.Install)}
	cmd := exec.Command("docker", cmdArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println(out.String())
	return nil
}
