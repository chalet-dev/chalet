package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func CheckDockerContainerExists(config *Config) error {
	containerName := fmt.Sprintf("chalet-%s", config.Name)

	// Check if the container exists
	existsCmd := exec.Command("docker", "ps", "-a", "--filter", fmt.Sprintf("name=^%s$", containerName), "--format", "{{.Names}}")
	var existsOut bytes.Buffer
	existsCmd.Stdout = &existsOut
	err := existsCmd.Run()
	if err != nil {
		fmt.Println("Error running Docker command:", err)
		return err
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
			fmt.Println("Error running Docker command:", err)
			return err
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
	fmt.Println("Creating container...")
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if config.ExposedPort == "" {
		config.ExposedPort = "7300"
	}

	cmd := exec.Command("docker",
		"create",
		"-it",
		"-v", fmt.Sprintf("%s:/app", cwd),
		"-p", fmt.Sprintf("%s:%s", config.ExposedPort, config.ServerPort),
		"--name", fmt.Sprintf("chalet-%s", config.Name),
		fmt.Sprintf("%s:%s", config.Lang, config.Version))

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

func startContainer(config *Config) error {
	fmt.Println("Starting container...")
	cmd := exec.Command("docker", "start", fmt.Sprintf("chalet-%s", config.Name))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func StopContainer(config *Config) error {
	fmt.Println("Stopping container...")
	cmd := exec.Command("docker", "stop", fmt.Sprintf("chalet-%s", config.Name))
	return cmd.Run()
}
