package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var errBadCh = errors.New("bad character")
var errOutOfRange = errors.New("out of range")

// function search slashed symbol with setted rules.

func parseSlash(pos int, inputStr map[int]rune) (string, error) {
	if pos >= len(inputStr) {
		return "", errOutOfRange
	}

	ch := string(inputStr[pos])

	if ch == `\` || (ch >= `0` && ch <= `9`) {
		return ch, nil
	}
	return "", errBadCh
}

// unpack difficult string with slash.

func unpackSlash(inputStr map[int]rune) (string, error) {
	var outputStr strings.Builder

	i := 0
	j := 1
	count := 0
	var slashedch string

	for i < len(inputStr) {
		// classic.
		count = 1
		if _, err := strconv.Atoi(string(inputStr[i])); err == nil {
			return "", ErrInvalidString
			// if founded `\``.
		} else if string(inputStr[i]) == `\` {
			if _slashedch, errf := parseSlash(i+1, inputStr); errf == nil {
				slashedch = _slashedch // save slashed symb.
				j++                    // move j to next position.
				i += 3                 // move i to next position of character.
			} else if errors.Is(errf, errBadCh) {
				return "", ErrInvalidString
			}
		}

		// search repeat count.
		if j < len(inputStr) {
			if _count, err := strconv.Atoi(string(inputStr[j])); err == nil {
				count = _count
				j += 2
			} else {
				j++
				// move i back.
				if len(slashedch) > 0 {
					i--
				}
			}
		}

		// make output string.
		if len(slashedch) > 0 {
			outputStr.WriteString(strings.Repeat(slashedch, count))
			slashedch = ""
		} else {
			outputStr.WriteString(strings.Repeat(string(inputStr[i]), count))
			if count != 1 {
				i += 2
			} else {
				i++
			}
		}
	}
	return outputStr.String(), nil
}

func unpack(inputStr map[int]rune) (string, error) {
	var outputStr strings.Builder

	i := 0
	j := 1

	for i < len(inputStr) {
		if _, err := strconv.Atoi(string(inputStr[i])); err == nil {
			return "", ErrInvalidString
		} else if j < len(inputStr) {
			if repeatCount, err := strconv.Atoi(string(inputStr[j])); err == nil {
				outputStr.WriteString(strings.Repeat(string(inputStr[i]), repeatCount))
				i += 2 // move i to next character.
				j += 2 // move j to next repeat-character if this character will be finded.
				continue
			}
		}
		// make output string.
		outputStr.WriteRune(inputStr[i])
		i++
		j++
	}
	return outputStr.String(), nil
}

func Unpack(inputStr string) (string, error) {
	var runeMap = make(map[int]rune)
	i := 0

	// make aligned hash of characters, if character implemented has non-standard rune.
	// etc. emoji, domino;).
	for _, value := range inputStr {
		runeMap[i] = value
		i++
	}

	// two scenarios.
	if strings.Contains(inputStr, `\`) {
		return unpackSlash(runeMap)
	}
	return unpack(runeMap)
}
