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
	data := []rune(s)
	lastValidIndex := len(data) - 1

	for i, r := range data {
		if unicode.IsDigit(r) {
			continue
		}

		if (i < lastValidIndex) && (unicode.IsDigit(data[i+1])) {
			n, err := strconv.Atoi(string(data[i+1]))
			if err != nil {
				return "", err
			}

			builder.WriteString(strings.Repeat(string(r), n))
			continue
		}

		builder.WriteString(string(r))
	}

	return builder.String(), nil
}

func isNotValidString(s string) bool {
	return regexp.MustCompile(`^\d|\d\d| `).MatchString(s)
}
