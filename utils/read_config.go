package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type Command struct {
	Run string `yaml:"run"`
}

type Config struct {
	Name           string            `yaml:"name"`
	Lang           string            `yaml:"lang"`
	Version        string            `yaml:"version"`
	ServerPort     string            `yaml:"server_port"`
	ExposedPort    string            `yaml:"exposed_port"`
	Commands       Command           `yaml:"commands"`
	CustomCommands map[string]string `yaml:"custom_commands"`
}

func ReadConfig() (*Config, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return nil, err
	}

	// Open the YAML file
	file, err := os.Open(cwd + "/chalet.yaml")
	if err != nil {
		fmt.Println("Error opening YAML file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the YAML file
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return nil, err
	}

	return &config, nil
}
