package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func checkDigit(current rune, prev *string, isSlash *bool) (bool, error) {
	if *prev == "" {
		return false, ErrInvalidString
	} else if *isSlash {
		*prev = string(current)
		*isSlash = false
		return false, nil
	}
	return true, nil
}

func Unpack(inputStr string) (string, error) {
	var outputStr strings.Builder
	prev := ""
	isSlash := false
	for _, current := range inputStr {
		switch {
		// if current character is digit -> check digit and write character n-count.
		case unicode.IsDigit(current):
			if status, err := checkDigit(current, &prev, &isSlash); errors.Is(err, ErrInvalidString) {
				return "", ErrInvalidString
			} else if status {
				outputStr.WriteString(strings.Repeat(prev, int(current-'0')))
				prev = ""
				isSlash = false
				continue
			}
		// if current character non-digit, write previous or make slashed.
		case prev != "":
			// make slashed character.
			if isSlash {
				if string(current) != `\` {
					return "", ErrInvalidString
				}
				isSlash = false
				prev = string(current)
				continue
			}
			outputStr.WriteString(prev)
		}
		// clear previous and save current in previous.
		prev = string(current)
		isSlash = string(current) == `\`
	}
	outputStr.WriteString(prev)
	return outputStr.String(), nil
}
