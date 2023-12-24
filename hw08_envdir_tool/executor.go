package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, variable := range env {
		if variable.NeedRemove {
			os.Unsetenv(name)
			continue
		}
		os.Setenv(name, variable.Value)
	}
	command, params := cmd[0], cmd[1:]
	executor := exec.Command(command, params...)
	executor.Env = os.Environ()
	executor.Stdin = os.Stdin
	executor.Stdout = os.Stdout
	executor.Stderr = os.Stderr

	executor.Start()
	if err := executor.Wait(); err != nil {
		return 1
	}
	return 0
}
