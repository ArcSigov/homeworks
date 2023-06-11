package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input_str string) (string, error) {
	var rune_map = make(map[int]rune)
	var output_str strings.Builder

	i := 0 // первый курсор отслеживает нахождение символа (не числа!) в строке
	j := 1 // второй курсор отслеживает нахождение числа в строке
	// с помощью которого нужно выполнить нужное количество повторений символа

	//первый шаг - формирование хэш-таблицы строки
	//при содержании в строке различных символов unicode его размер может быть разный
	//для этого выполняется выравнивание индексов
	for _, value := range input_str {
		rune_map[i] = value
		i++
	}
	i = 0

	//проход по хэш-таблице и распаковка строки
	for i < len(rune_map) {
		if _, err := strconv.Atoi(string(rune_map[i])); err == nil {
			return "", ErrInvalidString
		} else if j < len(rune_map) {
			if repeat_count, err := strconv.Atoi(string(rune_map[j])); err == nil {
				output_str.WriteString(strings.Repeat(string(rune_map[i]), repeat_count))
				i += 2
				j += 2
				continue
			}
		}
		output_str.WriteRune(rune_map[i])
		i++
		j++
	}
	return output_str.String(), nil
}
