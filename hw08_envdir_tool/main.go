package main

import (
	"log/slog"
	"os"
)

func main() {
	args := os.Args
	dir := args[1]
	cmd := args[2:]
	env, err := ReadDir(dir)
	if err != nil {
		slog.Error("error while reading env in dir", "dir", dir, "error", err)
		return
	}
	resultCode := RunCmd(cmd, env)

	os.Exit(resultCode)
}
