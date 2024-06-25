package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func Execute(name string, command string) error {
	cmdArgs := []string{"exec", fmt.Sprintf("chalet-%s", name), "sh", "-c", fmt.Sprintf("cd /chalet && %s", command)}
	cmd := exec.Command("docker", cmdArgs...)

	done := make(chan error, 1)

	if err := createLogPipes(cmd, done); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Println("Stopping the container...")

	stopErr := StopContainer(name)
	if stopErr != nil {
		fmt.Println("Failed to stop container:", stopErr)
		return stopErr
	}
	fmt.Println("Chalet stopped successfully.")
	return fmt.Errorf("process interrupted by signal: %v", sig)
}


func createLogPipes(cmd *exec.Cmd, done chan error) error {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating stderr pipe: %v", err)
	}

	go func() {
		done <- cmd.Run()
	}()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	return nil
}