package anagram

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func wordsSet(words []string) []string {
	set := make(map[string]bool)

	for _, wrd := range words {
		set[wrd] = true
	}

	result := make([]string, 0, len(set))

	for wrd := range set {
		result = append(result, wrd)
	}

	return result
}

func symbols(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//SearchAnagram ищет все множества анаграмм по словарю
func SearchAnagram(strList []string) (map[string][]string, error) {
	if len(strList) == 0 {
		return nil, fmt.Errorf("некорректные данные")
	}
	sets := make(map[string][]string)

	//ищем все анаграммы
	for _, str := range strList {
		//если 1 символ-пропускаем
		if utf8.RuneCountInString(str) == 1 {
			continue
		}
		str = strings.ToLower(str)
		anagramSet := symbols(str) //сортируем символы для определения множества
		sets[anagramSet] = append(sets[anagramSet], str)
	}

	//создаем мапу нужного вида
	anagrams := make(map[string][]string, len(sets))
	for anagram := range sets {
		firstEl := sets[anagram][0]
		//убираем повторения и сортируем слова
		anagrams[firstEl] = append(anagrams[firstEl], wordsSet(sets[anagram])...)
		sort.Strings(anagrams[firstEl])
	}

	return anagrams, nil

}
