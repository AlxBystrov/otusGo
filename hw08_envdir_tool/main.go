package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	dir := args[1]
	cmd := args[2:]
	env, err := ReadDir(dir)
	if err != nil {
		fmt.Printf("error while reading env in dir %s: %s", dir, err)
		return
	}
	resultCode := RunCmd(cmd, env)

	os.Exit(resultCode)
}
