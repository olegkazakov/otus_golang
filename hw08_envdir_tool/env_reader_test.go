package main

import (
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

const (
	validDirPath             = "testdata/env"
	emptyDirPath             = "testdata/empty"
	invalidDirectoryPath     = "testdata/env2/"
	incorrectFileNameDirPath = "testdata/env3"
	incorrectFileNamePath    = "testdata/env3/BAR="
	fileAsDirectoryPath      = "testdata/env4/"
	fileAsDirectoryFilePath  = "testdata/env4/BAR/"
)

func TestReadDir(t *testing.T) {
	defer func() {
		os.Remove(emptyDirPath)
		os.RemoveAll(invalidDirectoryPath)
		os.RemoveAll(incorrectFileNameDirPath)
		os.RemoveAll(fileAsDirectoryPath)
	}()

	t.Run("valid case", func(t *testing.T) {
		expectedRes := Environment{
			"BAR": EnvValue{
				Value: "bar",
			},
			"EMPTY": EnvValue{
				NeedRemove: true,
			},
			"FOO": EnvValue{
				Value: "   foo\nwith new line",
			},
			"HELLO": EnvValue{
				Value: `"hello"`,
			},
			"UNSET": EnvValue{
				NeedRemove: true,
			},
		}

		res, err := ReadDir(validDirPath)

		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("empty directory", func(t *testing.T) {
		err := os.Mkdir(emptyDirPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		res, err := ReadDir(emptyDirPath)
		require.NoError(t, err)
		require.Len(t, res, 0)
	})

	t.Run("reading directory error", func(t *testing.T) {
		err := os.Mkdir(invalidDirectoryPath, 0000)
		if err != nil {
			log.Fatal(err)
		}

		res, err := ReadDir(invalidDirectoryPath)

		require.Error(t, err)
		require.EqualError(t, err, ErrReadDirectory.Error())
		require.Len(t, res, 0)
	})

	t.Run("incorrect file name error", func(t *testing.T) {
		err := os.Mkdir(incorrectFileNameDirPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(incorrectFileNamePath)
		if err != nil {
			log.Fatal(err)
		}

		res, err := ReadDir(incorrectFileNameDirPath)

		require.Error(t, err)
		require.EqualError(t, err, ErrIncorrectFileName.Error())
		require.Len(t, res, 0)
	})

	t.Run("file is a directory error", func(t *testing.T) {
		err := os.Mkdir(fileAsDirectoryPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Mkdir(fileAsDirectoryFilePath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		res, err := ReadDir(fileAsDirectoryPath)

		require.Error(t, err)
		require.EqualError(t, err, ErrFileIsDirectory.Error())
		require.Len(t, res, 0)
	})
}
