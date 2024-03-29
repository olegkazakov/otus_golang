package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		name     string
	}{
		{name: "success case", input: "a4bc2d5e", expected: "aaaabccddddde"},
		{name: "unpack without digits", input: "abccd", expected: "abccd"},
		{name: "empty string", input: "", expected: ""},
		{name: "zero digit", input: "aaa0b", expected: "aab"},
		{name: "one digit", input: "ac1b", expected: "acb"},
		{name: "line break symbol", input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{name: "тест кириллицы", input: "тес3тКириллиц2ы2", expected: "тессстКириллиццыы"},
		{name: "тест кириллицы 2", input: "тес1тКириллицц0ы", expected: "тестКириллицы"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", " abcd", "abc def", "qwerty   "}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
