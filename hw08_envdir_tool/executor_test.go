package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testFile = "test_file.sh"
)

func TestRunCmd(t *testing.T) {
	t.Run("no command", func(t *testing.T) {
		code := RunCmd([]string{}, Environment{})
		require.Equal(t, ExitCodeError, code)
	})

	t.Run("no args", func(t *testing.T) {
		err := ioutil.WriteFile(testFile, []byte("#!/usr/bin/env bash\ngrep"), 0644)
		if err != nil {
			fmt.Println(err)
		}

		t.Cleanup(remove)
		args := []string{"bash", testFile}

		code := RunCmd(args, Environment{})
		require.Equal(t, 2, code)
	})
}

func remove() {
	if err := os.Remove(testFile); err != nil {
		fmt.Println(err)
	}
}
