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

func Unpack(inputStr string) (string, error) {
	var outputStr strings.Builder
	prev := characterParam{"", false, false}
	start := true
	for _, current := range inputStr {
		if start {
			prev.character = string(current)
			if prev.character == `\` {
				prev.isSlash = true
			} else if prev.character >= "0" && prev.character <= "9" {
				return "", ErrInvalidString
			}
			start = false
			continue
		}
		repeatCount, err := strconv.Atoi(string(current))
		if err == nil {
			switch {
			case !prev.isSlashed && (prev.character == "" || prev.character >= "0" && prev.character <= "9"):
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
			// work with characters.
		} else {
			switch {
			// if prevcharacter empty -> save current.
			case prev.character == "":
				prev.character = string(current)
			// if prevcharacter is slash -> check rules for slash.
			case prev.isSlash:
				if string(current) != `\` && !(string(current) >= "0" && string(current) <= "9") {
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
	}
	// print last character.
	if prev.character != "" {
		outputStr.WriteString(prev.character)
	}
	return outputStr.String(), nil
}
