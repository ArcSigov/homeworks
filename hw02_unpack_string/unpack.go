package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

type characterParam struct {
	character string
	isSlash   bool
	isSlashed bool
}

func isDigit(ch string) bool {
	return ch >= "0" && ch <= "9"
}

func prevInit(inputStr string, prev *characterParam) error {
	r := []rune(inputStr)
	if len(r) > 0 {
		prev.character = string(r[0])
	}
	if prev.character == `\` {
		prev.isSlash = true
	} else if isDigit(prev.character) {
		return ErrInvalidString
	}
	return nil
}

func Unpack(inputStr string) (string, error) {
	var outputStr strings.Builder
	prev := characterParam{"", false, false}
	if err := prevInit(inputStr, &prev); err != nil {
		return "", ErrInvalidString
	}
	start := true
	for _, current := range inputStr {
		if start {
			start = false
			continue
		}
		if repeatCount, err := strconv.Atoi(string(current)); err == nil {
			switch {
			case !prev.isSlashed && (prev.character == "" || isDigit(prev.character)):
				return "", ErrInvalidString
			// formatting slash.
			case prev.isSlash:
				prev.character = string(current)
				prev.isSlash = false
				prev.isSlashed = true
			// print formatted character repeat count.
			default:
				outputStr.WriteString(strings.Repeat(prev.character, repeatCount))
				prev.character = ""
				prev.isSlash = false
				prev.isSlashed = false
			}
			continue
		}
		// work with characters.
		switch {
		// if prevcharacter empty -> save current.
		case prev.character == "":
			prev.character = string(current)
		// if prevcharacter is slash -> check rules for slash.
		case prev.isSlash:
			if string(current) != `\` && !isDigit(string(current)) {
				return "", ErrInvalidString
			}
			prev.isSlash = false
			prev.isSlashed = true
			prev.character = string(current)
		// print non-formatted characters.
		default:
			outputStr.WriteString(prev.character)
			prev.character = string(current)
			if prev.character == `\` {
				prev.isSlash = true
			}
		}
	}
	// print last character.
	outputStr.WriteString(prev.character)
	return outputStr.String(), nil
}
