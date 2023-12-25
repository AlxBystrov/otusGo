package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	myEnv := make(Environment)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		if strings.Contains(info.Name(), "=") {
			continue
		}
		if info.Size() == 0 {
			myEnv[info.Name()] = EnvValue{Value: "", NeedRemove: true}
			continue
		}
		source, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		sc := bufio.NewScanner(source)
		readed := sc.Scan()
		if !readed {
			continue
		}
		value := sc.Text()
		value = strings.TrimRight(value, " \t")
		value = string(bytes.ReplaceAll([]byte(value), []byte("\x00"), []byte("\n")))

		myEnv[info.Name()] = EnvValue{Value: value, NeedRemove: false}
	}

	return myEnv, nil
}
