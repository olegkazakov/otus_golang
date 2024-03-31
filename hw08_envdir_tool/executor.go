package main

import (
	"errors"
	"os"
	"os/exec"
)

const (
	exitCodeOk    = 0
	exitCodeError = 127
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, envVar := range env {
		if _, ok := os.LookupEnv(key); ok {
			_ = os.Unsetenv(key)
		}

		if !envVar.NeedRemove {
			_ = os.Setenv(key, envVar.Value)
			continue
		}
	}

	executedCmd := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	executedCmd.Stdout = os.Stdout
	executedCmd.Stdin = os.Stdin
	executedCmd.Stderr = os.Stderr

	if err := executedCmd.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}

	return exitCodeOk
}
