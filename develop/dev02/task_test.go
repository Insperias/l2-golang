package unpack

import (
	"testing"
	"unicode"
)

func TestIsDigit(t *testing.T) {
	validDigitRunes := "012345679"
	for _, digitRune := range validDigitRunes {
		if !unicode.IsDigit(digitRune) {
			t.Errorf("Символ \"%s\" [rune %d] не является числом", string(digitRune), digitRune)
		}
	}

	invalidDigitRunes := "abXY@%"
	for _, digitRune := range invalidDigitRunes {
		if unicode.IsDigit(digitRune) {
			t.Errorf("Символ \"%s\" [rune %d] является числом", string(digitRune), digitRune)
		}
	}
}

func TestIsEscapeSymbol(t *testing.T) {
	validEscapeRunes := `\`
	for _, escapeRune := range validEscapeRunes {
		if !isEscapeSymbol(escapeRune) {
			t.Errorf("Символ \"%s\" [rune %d] не является escape-символом", string(escapeRune), escapeRune)
		}
	}

	invalidEscapeRunes := `|/@!`
	for _, escapeRune := range invalidEscapeRunes {
		if isEscapeSymbol(escapeRune) {
			t.Errorf("Символ \"%s\" [rune %d] является escape-символом", string(escapeRune), escapeRune)

		}
	}
}

func TestRegular(t *testing.T) {
	data := "a4b3c2d6e"
	expect := "aaaabbbccdddddde"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)
	}
}

func TestSimple(t *testing.T) {
	data := "abcde"
	expect := "abcde"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)
	}
}

func TestOnlyNumbers(t *testing.T) {
	data := "45"

	if result, err := Unpack(data); err == nil {
		t.Errorf("Пропущены неверные данные с результатом \"%s\", ожидалась ошибка", result)
	}
}

func TestEmptyString(t *testing.T) {
	if result, _ := Unpack(""); result != "" {
		t.Errorf("Неверные данные: ожидалась пустая строка, получена \"%s\"", result)
	}
}

func TestEscapedNumber(t *testing.T) {
	data := `abc\1\2`
	expect := "abc12"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)

	}
}

func TestPackedNumber(t *testing.T) {
	data := `abc\13`
	expect := `abc111`

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)

	}
}

func TestPackedSlash(t *testing.T) {
	data := `abc\\4`
	expect := `abc\\\\`

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)

	}
}

func TestComplexSymbols(t *testing.T) {
	data := "€3"
	expect := "€€€"

	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)

	}
}

func TestMoreThenNineNumber(t *testing.T) {
	data := "a11b"
	expect := "aaaaaaaaaaab"
	if result, _ := Unpack(data); result != expect {
		t.Errorf("Неверный результат: ожидалось \"%s\", получено \"%s\"", expect, result)

	}
}
