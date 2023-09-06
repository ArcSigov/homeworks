package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func prepareInput(fromPath string, offset int64, limit *int64) (*os.File, error) {
	inputFile, err := os.Open(fromPath)
	if err != nil {
		return nil, ErrUnsupportedFile
	}
	info, _ := inputFile.Stat()
	if info.Size() <= offset {
		return nil, ErrOffsetExceedsFileSize
	}
	inputFile.Seek(offset, io.SeekStart)

	if *limit == 0 {
		*limit = info.Size()
	}
	return inputFile, nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := prepareInput(fromPath, offset, &limit)
	if err != nil {
		return err
	}
	defer file.Close()
	outputFile, _ := os.Create(toPath)
	defer outputFile.Close()

	buf := make([]byte, 1)
	copied := int64(0)
	for copied < limit {
		count, err := file.Read(buf)
		copied += int64(count)
		if err == io.EOF {
			break
		}
		outputFile.Write(buf)
	}
	return nil
}
