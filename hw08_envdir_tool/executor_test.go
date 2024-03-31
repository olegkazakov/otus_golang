package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env := Environment{
		"BAR":   {"bar", false},
		"UNSET": {"", true},
		"EMPTY": {"", false},
		"FOO":   {"   foo\nwith new line", false},
		"HELLO": {"\"hello\"", false},
	}

	pwd, _ := os.Getwd()

	t.Run("valid case", func(t *testing.T) {
		cmd := []string{
			"/bin/bash",
			pwd + "/testdata/echo.sh",
			"arg1=1",
			"arg2=2",
		}

		returnCode := RunCmd(cmd, env)
		require.Equal(t, exitCodeOk, returnCode)
	})

	t.Run("error", func(t *testing.T) {
		cmd := []string{
			"/bin/bash",
			pwd + "/testdata/echo1.sh",
		}

		returnCode := RunCmd(cmd, env)
		require.Equal(t, exitCodeError, returnCode)
	})
}
