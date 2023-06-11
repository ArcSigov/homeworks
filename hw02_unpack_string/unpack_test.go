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
		//добавление рун для проверки
		{input: "🁐5", expected: "🁐🁐🁐🁐🁐"},
		{input: "㌀2🂲3", expected: "㌀㌀🂲🂲🂲"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
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
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

// тест проверки паник при работе с функцией Unpack
// при неправильном алгоритме возможна ситуация с формированием полного числа
// прежде чем выбить ошибку формата
func TestPanic(t *testing.T) {
	input_str := "a4444444444444444444444444444444444444444b99999999999999999999999999с88888"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function Unpack call panic!")
		}
	}()
	Unpack(input_str)
}
