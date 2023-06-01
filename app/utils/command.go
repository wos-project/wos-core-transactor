package utils

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExecuteCommand(cmd string, timeoutSeconds int, args ...string) (string, error) {

	context, _ := context.WithTimeout(context.Background(), time.Second * time.Duration(timeoutSeconds))

	command := exec.CommandContext(context, cmd, args...)

	output, err := command.Output()
	if err != nil {
		return "", err
	}

	if context.Err() != nil {
		return "", fmt.Errorf("timeout")
	}

	return string(output), nil
}