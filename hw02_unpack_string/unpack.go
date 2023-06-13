package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type characterParam struct {
	character string
	isSlash   bool
}

func checkDigit(current rune, prev *characterParam) (bool, error) {
	if prev.character == "" {
		return false, ErrInvalidString
	} else if prev.isSlash {
		prev.character = string(current)
		prev.isSlash = false
		return false, nil
	}
	return true, nil
}

func clearPrev(prev *characterParam) {
	prev.character = ""
	prev.isSlash = false
}

func Unpack(inputStr string) (string, error) {
	var outputStr strings.Builder
	var prev characterParam
	for _, current := range inputStr {
		switch {
		// if current character is digit -> check digit and write character n-count.
		case unicode.IsDigit(current):
			if status, err := checkDigit(current, &prev); errors.Is(err, ErrInvalidString) {
				return "", ErrInvalidString
			} else if status {
				outputStr.WriteString(strings.Repeat(prev.character, int(current-'0')))
				clearPrev(&prev)
				continue
			}
		// if current character non-digit, write previous or make slashed.
		case prev.character != "":
			// make slashed character.
			if prev.isSlash {
				if string(current) != `\` {
					return "", ErrInvalidString
				}
				prev.isSlash = false
				prev.character = string(current)
				continue
			}
			outputStr.WriteString(prev.character)
		}
		// clear previous and save current in previous.
		prev.character = string(current)
		prev.isSlash = string(current) == `\`
	}
	outputStr.WriteString(prev.character)
	return outputStr.String(), nil
}
