package main

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrReadDirectory     = errors.New("reading directory error")
	ErrFileIsDirectory   = errors.New("file is a directory")
	ErrIncorrectFileName = errors.New("file name could not contain '=' symbol")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	filesData, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrReadDirectory
	}

	envVars := make(Environment, len(filesData))
	for _, fileInfo := range filesData {
		envVar, err := ReadVarFromFile(dir, fileInfo)
		if err != nil {
			return nil, err
		}
		envVars[fileInfo.Name()] = *envVar
	}

	// Place your code here
	return envVars, nil
}

func ReadVarFromFile(dir string, fileInfo fs.DirEntry) (*EnvValue, error) {
	if fileInfo.IsDir() {
		return nil, ErrFileIsDirectory
	}

	if strings.Contains(fileInfo.Name(), "=") {
		return nil, ErrIncorrectFileName
	}

	file, err := os.Open(filepath.Join(dir, fileInfo.Name()))
	if err != nil {
		return nil, err
	}

	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	value := filterEnvValue(scanner.Text())
	return &EnvValue{
		Value:      value,
		NeedRemove: len(value) == 0,
	}, nil
}

func filterEnvValue(s string) string {
	value := strings.TrimRightFunc(s, unicode.IsSpace)
	return strings.ReplaceAll(value, "\x00", "\n")
}
