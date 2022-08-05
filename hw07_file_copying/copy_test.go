package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	fromPath := "testdata/input.txt"
	toPath := "output.txt"

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", toPath, 0, 0)
		require.Truef(t, errors.Is(err, ErrUnsupportedFile), "%v", err)
	})

	t.Run("offset exceeds fileSize", func(t *testing.T) {
		err := Copy(fromPath, toPath, 10000, 0)
		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "%v", err)
	})

	t.Run("limit is negative", func(t *testing.T) {
		err := Copy(fromPath, toPath, 0, -1)
		require.Truef(t, errors.Is(err, ErrNegativeLimit), "%v", err)
	})

	t.Run("no such file or directory", func(t *testing.T) {
		err := Copy("testdat/input.txt", toPath, 0, 0)
		require.Truef(t, os.IsNotExist(err), "%v", err)
	})
}
