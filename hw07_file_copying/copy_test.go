package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	var (
		source      = "testdata/input.txt"
		destination = "testdata/output.txt"
		offset      int64
		limit       int64
	)

	t.Run("empty source filename", func(t *testing.T) {
		err := Copy("", destination, offset, limit)
		require.ErrorIs(t, err, ErrSourceFileIsNotSpecified)
	})

	t.Run("empty destination filename", func(t *testing.T) {
		err := Copy(source, "", offset, limit)
		require.ErrorIs(t, err, ErrDestinationFileIsNotSpecified)
	})

	t.Run("source and destination file paths are identical", func(t *testing.T) {
		err := Copy(source, source, offset, limit)
		require.ErrorIs(t, err, ErrSourceAndDestinationPathsAreIdentical)
	})

	t.Run("unsupported source file", func(t *testing.T) {
		err := Copy("/dev/urandom", destination, offset, limit)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("negative offset", func(t *testing.T) {
		err := Copy(source, destination, -1, limit)
		require.ErrorIs(t, err, ErrNegativeOffset)
	})

	t.Run("negative limit", func(t *testing.T) {
		err := Copy(source, destination, offset, -1)
		require.ErrorIs(t, err, ErrNegativeLimit)
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		sourceFileStat, _ := os.Stat(source)
		err := Copy(source, destination, sourceFileStat.Size()+1, limit)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("limit greater than source file size", func(t *testing.T) {
		sourceFileStat, _ := os.Stat(source)
		sourceFileSize := sourceFileStat.Size()
		err := Copy(source, destination, offset, sourceFileSize+1024)
		require.Nil(t, err)
		destinationFileStat, _ := os.Stat(destination)
		require.Equal(t, sourceFileSize, destinationFileStat.Size())
	})

	t.Run("successful case", func(t *testing.T) {
		sourceFileStat, _ := os.Stat(source)
		err := Copy(source, destination, 0, 0)
		require.Nil(t, err)
		destinationFileStat, _ := os.Stat(destination)
		require.Equal(t, sourceFileStat.Size(), destinationFileStat.Size())
	})
}
