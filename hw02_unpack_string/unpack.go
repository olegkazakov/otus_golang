package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var re = regexp.MustCompile(`^\d|\d\d| `)

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	if re.MatchString(s) {
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
