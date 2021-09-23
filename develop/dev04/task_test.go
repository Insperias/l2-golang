package anagram

import (
	"testing"
)

func TestEmptyArray(t *testing.T) {
	var data []string
	if result, err := SearchAnagram(data); err == nil {
		t.Errorf("Пропущены неверные данные с результатом \"%s\", ожидалась ошибка", result)
	}
}

func TestSimpleSymbol(t *testing.T) {
	data := []string{"а", "а", "в", "б"}

	if result, _ := SearchAnagram(data); len(result) != 0 {
		t.Errorf("Неверный результат: ожидалась длина 0, получена длина %d", len(result))
	}
}
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func comapareMap(map1, map2 map[string][]string) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key1, slice1 := range map1 {
		compare := false
		for key2, slice2 := range map2 {
			if key1 == key2 && equalSlice(slice1, slice2) {
				compare = true
				break
			}
		}
		if !compare {
			return false
		}
	}
	return true
}
func TestRegular(t *testing.T) {
	data := []string{"мама", "амам", "ммаа", "крик", "ккир", "кирк", "приморск"}
	expect := map[string][]string{
		"мама":     []string{"амам", "мама", "ммаа"},
		"крик":     []string{"кирк", "ккир", "крик"},
		"приморск": []string{"приморск"},
	}

	if result, _ := SearchAnagram(data); !comapareMap(result, expect) {
		t.Errorf("Неверный результат: ожидалось %v\n, получено %v\n", expect, result)
	}
}

func TestRepeatWords(t *testing.T) {
	data := []string{"мама", "мама", "папа", "папа", "амам", "апап"}
	expect := map[string][]string{
		"мама": []string{"амам", "мама"},
		"папа": []string{"апап", "папа"},
	}
	if result, _ := SearchAnagram(data); !comapareMap(result, expect) {
		t.Errorf("Неверный результат: ожидалось %v\n, получено %v\n", expect, result)
	}
}
