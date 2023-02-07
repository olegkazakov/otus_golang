package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	if isNotValidString(s) {
		return "", ErrInvalidString
	}

	var builder strings.Builder

	for i, r := range s {
		char := string(r)

		if unicode.IsDigit(r) {
			continue
		}

		if (i < len(s)-1) && (unicode.IsDigit(rune(s[i+1]))) {
			n, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				return "", err
			}

			builder.WriteString(strings.Repeat(char, n))
			continue
		}

		builder.WriteString(char)
	}

	return builder.String(), nil
}

func isNotValidString(s string) bool {
	return regexp.MustCompile(`^\d|\d\d| `).MatchString(s)
}
