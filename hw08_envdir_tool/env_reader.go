package main

import (
	"bufio"
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
	result, _ := os.ReadDir(dir)
	environmets := make(Environment)

	for _, file := range result {
		if file.IsDir() {
			continue
		}

		value, needRemove, err := readEnvFile(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		value = strings.ReplaceAll(value, "\x00", "\n")

		environmets[file.Name()] = EnvValue{Value: value, NeedRemove: needRemove}
	}

	return environmets, nil
}
func readEnvFile(path string) (string, bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		return scanner.Text(), false, nil
	}

	return "", false, nil
}
