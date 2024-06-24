package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CheckDockerContainerExists(containerName string) bool {
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
