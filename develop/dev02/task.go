package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

//Unpack распаковывает строки вида a4b3c2d => aaaabbbccd
func Unpack(data string) (string, error) {
	var buf string
	var isEscaped bool
	var result strings.Builder
	var count int

	for _, r := range data {
		s := string(r)

		if unicode.IsDigit(r) {
			//Если первый символ это цифра
			if buf == "" {
				return "", errors.New("некорректная строка")
			}

			//если это счетчик
			if buf != "" && !isEscaped {
				n, _ := strconv.Atoi(s)
				count = count*10 + n

				continue
			}
		}

		if buf != "" {
			//escape-символ
			if isEscapeSymbol(r) && !isEscaped {
				isEscaped = true
				continue
			}
			//сохраняем символ в результат
			isEscaped = false
			if count != 0 {
				result.WriteString(strings.Repeat(buf, count))
				count = 0
				buf = s
				continue
			} else {
				result.WriteString(buf)
				buf = s
				continue
			}
		}

		//добавляем символ в буфер
		buf = s
	}
	if count != 0 {
		result.WriteString(strings.Repeat(buf, count))
	} else {
		result.WriteString(buf)
	}
	return result.String(), nil
}

func isEscapeSymbol(r rune) bool {
	return r == 92
}
