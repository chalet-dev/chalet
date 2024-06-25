package utils

import (
	"errors"
	"fmt"
	"github.com/chalet/cli/logger"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func Execute(name string, command string) error {
	cmdArgs := []string{"exec", fmt.Sprintf("chalet-%s", name), "sh", "-c", fmt.Sprintf("cd /chalet && %s", command)}
	cmd := exec.Command("docker", cmdArgs...)

	done := make(chan error, 1)

	if err := CreateLogPipes(cmd, done); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	stopErr := StopContainer(name)
	if stopErr != nil {
		return errors.New("Failed to stop container: " + stopErr.Error())
	}
	logger.Info("Chalet stopped successfully.")
	return fmt.Errorf("process interrupted by signal: %v", sig)
}
