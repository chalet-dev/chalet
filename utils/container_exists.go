package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func CheckDockerContainerExists(config *Config) error {
	cmd := exec.Command("docker", "ps", "-a", "--filter", fmt.Sprintf("name=^%s$", config.Name), "--format", "{{.Names}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running Docker command:", err)
		return err
	}

	if !bytes.Contains(out.Bytes(), []byte(config.Name)) {
		err = createContainer(config)
		if err != nil {
			return err
		}
		err = startContainer(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func createContainer(config *Config) error {
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

func startContainer(config *Config) error {
	fmt.Println("Starting container...")
	cmd := exec.Command("docker", "start", config.Name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
