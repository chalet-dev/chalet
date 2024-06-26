/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"errors"
	"fmt"
	"github.com/chalet/cli/logger"
	"github.com/chalet/cli/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"strings"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new chalet project",
	Long: `Initializes a new chalet project by creating a chalet.yaml file,
	which will contain the configuration for the project. For example:
	chalet init -n project-name`,
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.Config{
			Name:       cmd.Flag("name").Value.String(),
			Lang:       cmd.Flag("language").Value.String(),
			Version:    cmd.Flag("version").Value.String(),
			ServerPort: cmd.Flag("port").Value.String(),
			Commands: utils.Command{
				Run: cmd.Flag("run").Value.String(),
			},
		}
		err := initProject(config)
		if err != nil {
			logger.Error(err.Error())
		}
	},
}

type configForMarshalling struct {
	Name           string            `yaml:"name"`
	Lang           string            `yaml:"lang"`
	Version        string            `yaml:"version"`
	ServerPort     string            `yaml:"server_port"`
	Commands       utils.Command     `yaml:"commands"`
	CustomCommands map[string]string `yaml:"custom_commands"`
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().StringP("name", "n", "", "container project name")
	initCmd.Flags().StringP("language", "l", "", "project language")
	initCmd.Flags().StringP("version", "", "", "project version")
	initCmd.Flags().StringP("port", "p", "", "server port")
	initCmd.Flags().StringP("run", "r", "", "run command")
}

func initProject(config utils.Config) error {
	cmd := exec.Command("docker", "version")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return errors.New("docker is not installed or not running")
	}

	if !strings.Contains(string(output), "Server: Docker") {
		return errors.New("docker is not installed or not running")
	}

	// Marshal the struct to YAML
	yamlData, err := yaml.Marshal(&configForMarshalling{
		Name:           config.Name,
		Lang:           config.Lang,
		Version:        config.Version,
		ServerPort:     config.ServerPort,
		Commands:       config.Commands,
		CustomCommands: config.CustomCommands,
	})
	if err != nil {
		return errors.New(fmt.Sprint("error marshalling to YAML:", err))
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprint("error getting current directory:", err))
	}

	// Define the file path
	filePath := cwd + "/chalet.yml"

	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		logger.Warn("File already exists, skipping initialization")
		return nil
	} else if !os.IsNotExist(err) {
		return errors.New(fmt.Sprint("error checking file existence:", err))
	}

	// Open the YAML file for writing (create if not exists, truncate if exists)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New(fmt.Sprint("error opening chalet config file for writing:", err))
	}
	defer file.Close()

	// Write the YAML data to the file
	_, err = file.Write(yamlData)
	if err != nil {
		return errors.New(fmt.Sprint("error writing chalet config file:", err))
	}

	logger.Info("Chalet project initialized successfully")
	return nil
}
