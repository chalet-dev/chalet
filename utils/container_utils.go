package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/chalet/cli/logger"
	"os"
	"os/exec"
)

func CheckDockerContainerExists(config *Config) error {
	pullCmd := exec.Command("docker", "pull", fmt.Sprintf("%s:%s", config.Lang, config.Version))
	done := make(chan error, 1)
	if err := CreateLogPipes(pullCmd, done); err != nil {
		return err
	}
	containerName := fmt.Sprintf("chalet-%s", config.Name)

	// Check if the container exists
	existsCmd := exec.Command("docker", "ps", "-a", "--filter", fmt.Sprintf("name=^%s$", containerName), "--format", "{{.Names}}")
	var existsOut bytes.Buffer
	existsCmd.Stdout = &existsOut
	err := existsCmd.Run()
	if err != nil {
		return errors.New(fmt.Sprint("error running Docker command:", err))
	}

	if !bytes.Contains(existsOut.Bytes(), []byte(containerName)) {
		// If the container does not exist, create and start it
		err = createContainer(config)
		if err != nil {
			return err
		}
		err = startContainer(config)
		if err != nil {
			return err
		}
	} else {
		// Check if the container is running
		runningCmd := exec.Command("docker", "ps", "--filter", fmt.Sprintf("name=^%s$", containerName), "--format", "{{.Names}}")
		var runningOut bytes.Buffer
		runningCmd.Stdout = &runningOut
		err := runningCmd.Run()
		if err != nil {
			return errors.New(fmt.Sprint("error running Docker command:", err))
		}

		if !bytes.Contains(runningOut.Bytes(), []byte(containerName)) {
			// If the container exists but is not running, start it
			err = startContainer(config)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createContainer(config *Config) error {
	logger.Info("Creating container...")
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if config.ExposedPort == "" {
		config.ExposedPort = config.ServerPort
	}

	if config.Version == "" {
		config.Version = "latest"
	}

	cmd := exec.Command("docker",
		"create",
		"-it",
		"-v", fmt.Sprintf("%s:/chalet", cwd),
		"-p", fmt.Sprintf("%s:%s", config.ExposedPort, config.ServerPort),
		"--name", fmt.Sprintf("chalet-%s", config.Name),
		fmt.Sprintf("%s:%s", config.Lang, config.Version))

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func startContainer(config *Config) error {
	logger.Info("Starting container...")
	cmd := exec.Command("docker", "start", fmt.Sprintf("chalet-%s", config.Name))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func StopContainer(name string) error {
	logger.Info("Stopping container...")
	cmd := exec.Command("docker", "stop", fmt.Sprintf("chalet-%s", name))
	return cmd.Run()
}
