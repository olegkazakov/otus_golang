package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile                       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize                 = errors.New("offset exceeds file size")
	ErrSourceFileIsNotSpecified              = errors.New("source file is not specified")
	ErrDestinationFileIsNotSpecified         = errors.New("destination file is not specified")
	ErrSourceAndDestinationPathsAreIdentical = errors.New("source and destination paths are identical")
	ErrNegativeOffset                        = errors.New("negative offset")
	ErrNegativeLimit                         = errors.New("negative limit")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrSourceFileIsNotSpecified
	}

	if toPath == "" {
		return ErrDestinationFileIsNotSpecified
	}

	if fromPath == toPath {
		return ErrSourceAndDestinationPathsAreIdentical
	}

	if offset < 0 {
		return ErrNegativeOffset
	}

	if limit < 0 {
		return ErrNegativeLimit
	}

	sourceFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	destFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func() error {
		err = sourceFile.Close()
		if err != nil {
			return err
		}

		err = destFile.Close()
		if err != nil {
			return err
		}

		return nil
	}()

	sourceFileStat, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	sourceFileSize := sourceFileStat.Size()
	if offset > sourceFileSize {
		return ErrOffsetExceedsFileSize
	}

	readLimit := limit
	if readLimit == 0 {
		readLimit = sourceFileSize
	}

	if readLimit > sourceFileSize-offset {
		readLimit = sourceFileSize - offset
	}

	_, err = sourceFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	bar := pb.Full.Start64(readLimit)
	barReader := bar.NewProxyReader(sourceFile)
	defer bar.Finish()

	_, err = io.CopyN(destFile, barReader, readLimit)
	if err != nil {
		return err
	}

	return nil
}
