package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var badslashedch = errors.New("bad character")
var outOfRange = errors.New("out of range")

// function search slashed symbol with setted rules
func parseSlash(pos int, input_str map[int]rune) (string, error) {

	if pos >= len(input_str) {
		return "", outOfRange
	}

	ch := string(input_str[pos])

	if ch == `\` || (ch >= `0` && ch <= `9`) {
		return ch, nil
	} else {
		return "", badslashedch
	}

}

// unpack difficult string with slash
func unpackSlash(input_str map[int]rune) (string, error) {
	var output_str strings.Builder

	i := 0
	j := 1
	count := 0
	var slashedch string

	for i < len(input_str) {
		// classic
		count = 1
		if _, err := strconv.Atoi(string(input_str[i])); err == nil {
			return "", ErrInvalidString
			// if founded `\``
		} else if _slashedch, errf := parseSlash(i+1, input_str); errf == nil && string(input_str[i]) == `\` {
			slashedch = _slashedch // save slashed symb
			j++                    // move j to next position
			i += 3                 // move i to next position of character
		} else if string(input_str[i]) == `\` && errf == badslashedch {
			return "", ErrInvalidString
		}

		// search repeat count
		if j < len(input_str) {
			if _count, err := strconv.Atoi(string(input_str[j])); err == nil {
				count = _count
				j += 2
			} else {
				j++
				// move i back
				if len(slashedch) > 0 {
					i--
				}
			}
		}

		//make output string
		if len(slashedch) > 0 {
			output_str.WriteString(strings.Repeat(slashedch, count))
			slashedch = ""
		} else {
			output_str.WriteString(strings.Repeat(string(input_str[i]), count))
			if count != 1 {
				i += 2
			} else {
				i++
			}
		}
	}

	return output_str.String(), nil
}

func unpack(input_str map[int]rune) (string, error) {

	var output_str strings.Builder

	i := 0
	j := 1

	for i < len(input_str) {
		if _, err := strconv.Atoi(string(input_str[i])); err == nil {
			return "", ErrInvalidString
		} else if j < len(input_str) {
			if repeat_count, err := strconv.Atoi(string(input_str[j])); err == nil {
				output_str.WriteString(strings.Repeat(string(input_str[i]), repeat_count))
				i += 2 //move i to next character
				j += 2 //move j to next repeat-character if this character will be finded
				continue
			}
		}
		//make output string
		output_str.WriteRune(input_str[i])
		i++
		j++
	}
	return output_str.String(), nil
}

func Unpack(input_str string) (string, error) {

	var rune_map = make(map[int]rune)
	i := 0

	//make alligned hash of characters, if character implemented has non-standard rune
	// etc. emoji, domino;)
	for _, value := range input_str {
		rune_map[i] = value
		i++
	}

	// two scenarious
	if strings.Contains(input_str, `\`) {
		return unpackSlash(rune_map)
	} else {
		return unpack(rune_map)
	}
}
