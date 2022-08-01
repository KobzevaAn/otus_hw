package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
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
	var envs Environment = make(map[string]EnvValue)

	fsInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fInfo := range fsInfo {
		if fInfo.Size() == 0 {
			envs[fInfo.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		fPath := dir + "/" + fInfo.Name()
		fRead, err := os.Open(fPath)
		if err != nil {
			return nil, err
		}
		defer fRead.Close()

		buf := bufio.NewReader(fRead)
		firstLine, err := buf.ReadBytes('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
		}

		replaced := bytes.ReplaceAll(firstLine, []byte{0}, []byte("\n"))
		v := strings.TrimRight(string(replaced), " \t\n")

		envs[fInfo.Name()] = EnvValue{Value: v}
	}

	return envs, nil
}
