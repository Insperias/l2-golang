package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	After      int  // -A
	Before     int  // -B
	Context    int  // -C
	Count      bool // -c
	IgnoreCase bool // -i
	Invert     bool // -v
	Fixed      bool // -F
	LineNumber bool // -n
}

type checker interface {
	check(s string) bool
}

type regexChecker struct {
	R *regexp.Regexp
}

//NewRegexChecker ...
func NewRegexChecker(r *regexp.Regexp) *regexChecker {
	return &regexChecker{R: r}
}

func (c *regexChecker) check(s string) bool {
	return c.R.MatchString(s)
}

type equalChecker struct {
	pat string
}

//NewEqualChecker ...
func NewEqualChecker(pat string) *equalChecker {
	return &equalChecker{pat: pat}
}

func (c *equalChecker) check(s string) bool {
	return strings.Contains(s, c.pat)
}

//Grep ...
func Grep(r io.Reader, pattern string, fg *flags) (string, error) {
	var chr checker

	//если игнорировать регистр, то переводим все в нижний регистр
	if fg.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	if fg.Fixed {
		//если точное совпадение, то проверяем строку
		chr = NewEqualChecker(pattern)
	} else {
		//если паттерн, то компилируем как регулярное выражение
		rx, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		chr = NewRegexChecker(rx)
	}

	lm := fg.Context //количество строк перед найденной
	rm := fg.Context //количество строк после найденной
	if fg.Before > 0 {
		lm = fg.Before
	}
	if fg.After > 0 {
		rm = fg.After
	}

	outStrgsInd := make(map[int]bool) //карта для индексов строк, которые нужно будет вывести
	var resIds []int
	var allStr []string

	var counter, i int
	sc := bufio.NewScanner(r)
	//считываем строки
	for sc.Scan() {
		txt := sc.Text()
		allStr = append(allStr, txt)
		if fg.IgnoreCase {
			//если игнорировать регистр, то переводим в нижний регистр
			txt = strings.ToLower(txt)
		}

		//если нашли совпадение
		if chr.check(txt) {
			counter++
			//записываем все индексы строк перед найденной, если нужно
			if lm > 0 {
				for j := i - lm; j < i; j++ {
					if _, ok := outStrgsInd[j]; !ok {
						resIds = append(resIds, j)
						outStrgsInd[j] = false
					}
				}
			}

			//записываем все индексы строк после найденной, если нужно
			if rm > 0 {
				for j := i + rm; j > i; j-- {
					if _, ok := outStrgsInd[j]; !ok {
						resIds = append(resIds, j)
						outStrgsInd[j] = false
					}
				}
			}

			//записать индекс найденной строки и обозначить true, как найденную
			if _, ok := outStrgsInd[i]; !ok {
				resIds = append(resIds, i)
			}
			outStrgsInd[i] = true
		}
		i++

	}

	//если нужно только количество, то возвращаем и выходим
	if fg.Count {
		return fmt.Sprintf("%v matches", counter), nil
	}

	//сортируем индексы для упорядоченного вывода
	sort.Ints(resIds)

	var res []string
	lIndex := ""
	//если выводить найденные
	if !fg.Invert {
		match := ""
		for _, k := range resIds {
			//проверяем, чтобы не выйти за границы, т.к. могли быть добавлены элементы с индексами <0 и >len(allStr)
			if k > -1 && k < len(allStr) {
				//если это найденная строка, то добавить + как найденная
				if outStrgsInd[k] {
					match = "+ "
				} else {
					match = " "
				}
				if fg.LineNumber {
					//если нужно выводить с номером строки, то дописать
					lIndex = strconv.Itoa(k+1) + " "
				}
				//добавить все в результат
				res = append(res, fmt.Sprintf("%v%v%v", lIndex, match, allStr[k]))
			}
		}
	} else {
		//иначе нужно выводить те, которые не были найдены
		for i, v := range allStr {
			//если не добавляли в карту нужных для вывода строк
			if _, ok := outStrgsInd[i]; !ok {
				if fg.LineNumber {
					//добавить номер строки
					lIndex = strconv.Itoa(i+1) + " "
				}
				//добавить в результат
				res = append(res, fmt.Sprintf("%v%v", lIndex, v))
			}
		}
	}

	return strings.Join(res, "\n"), nil
}

func main() {
	fg := &flags{}

	flag.IntVar(&fg.After, "A", 0, "number of lines after match")
	flag.IntVar(&fg.Before, "B", 0, "number of lines before match")
	flag.IntVar(&fg.Context, "C", 0, "number of lines around match")
	flag.BoolVar(&fg.Count, "c", false, "count matches")
	flag.BoolVar(&fg.IgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&fg.Invert, "v", false, "exclude if match")
	flag.BoolVar(&fg.Fixed, "F", false, "concrete string, not a pattern")
	flag.BoolVar(&fg.LineNumber, "n", false, "print line number")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("grep: not enough arguments")
		return
	}

	//файл с исходными данными
	reader, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("grep err: ", err.Error())
		return
	}

	//паттерн
	pattern := flag.Arg(0)

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		fmt.Println("grep err: ", err.Error())
		return
	}

	fmt.Println(res)
}
