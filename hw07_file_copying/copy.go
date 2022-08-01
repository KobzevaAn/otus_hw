package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNegativeLimit         = errors.New("negative limit")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if limit < 0 {
		return ErrNegativeLimit
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	fileInfo, _ := fileFrom.Stat()
	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	if fileInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}
	fileFrom.Seek(offset, io.SeekStart)
	reader := io.Reader(fileFrom)

	bar := pb.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	if limit == 0 {
		limit = fileInfo.Size()
	}

	fileTo, _ := os.Create(toPath)
	defer fileTo.Close()
	writer := io.Writer(fileTo)

	if _, err := io.CopyN(writer, barReader, limit); err != nil {
		if errors.Is(err, io.EOF) {
			return err
		}
	}
	bar.Finish()

	return nil
}
