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
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// add runes.
		{input: "🁐5", expected: "🁐🁐🁐🁐🁐"},
		{input: "㌀2🂲3", expected: "㌀㌀🂲🂲🂲"},
		// add slashed.
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `\1qwe\4\5`, expected: `1qwe45`},
		{input: `\\1qwe\4\5`, expected: `\qwe45`},
		{input: `\\0qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `qwe\\3a4\`, expected: `qwe\\\aaaa\`},
		{input: `qwe\\3㌀2🂲3`, expected: `qwe\\\㌀㌀🂲🂲🂲`},
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
	invalidStrings := []string{"3abc", "45", "aaa10b", `qwe\\55de`, `qwe\nde`, `2qwe\\5`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

// panic test.
// this test check's call panic, if developer generate long repeat value with atoi function and string.
func TestPanic(t *testing.T) {
	inputStr := "a4444444444444444444444444444444444444444b99999999999999999999999999с88888"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function Unpack call panic!")
		}
	}()
	Unpack(inputStr)
}
