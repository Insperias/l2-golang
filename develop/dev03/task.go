package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flagSettings struct {
	column      int
	numeric     bool
	reverse     bool
	unduplicate bool
}

//читает данные из файла и возвращает []string всех строк
func readFile(fname string) ([]string, error) {
	strgs := make([]string, 0)
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("error: open file: err: [%s]", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		strgs = append(strgs, sc.Text())
	}

	return strgs, nil
}
func input() ([]string, error) {
	//список строк, т.к. сортировка может быть по столбцам
	strgs := make([]string, 0)
	if flag.Arg(0) == "-" {
		if flag.NArg() > 1 {
			//если после признака ввода из stdin переданы имена файлов, то нужно их прочесть
			i := 1
			for i < flag.NArg() {
				fstrgs, err := readFile(flag.Arg(i))
				if err != nil {
					return nil, err
				}
				strgs = append(strgs, fstrgs...)
				i++
			}
		}
		//ввод с клавиатуры (как и в man sort(1))
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			strgs = append(strgs, sc.Text())
		}

		if err := sc.Err(); err != nil {
			return nil, fmt.Errorf("error: reading stdio: err: [%s]", err)
		}
	} else { //иначе переданы пути к файлам с данными
		for _, arg := range flag.Args() {
			fstrgs, err := readFile(arg)
			if err != nil {
				return nil, err
			}
			strgs = append(strgs, fstrgs...)
		}
	}

	return strgs, nil
}

func removeDuplicate(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func stringSortColumn(strSlice [][]string, column int, reverse bool) [][]string {
	res := make([][]string, 0, len(strSlice))
	res = append(res, strSlice...)

	//если сортировать в обратном порядке
	if reverse {
		sort.Slice(res, func(i, j int) bool {
			//проверка на существование столбца
			if len(res[i]) < column {
				return false
			} else if len(res[j]) < column {
				return true
			}

			return res[i][column-1] > res[j][column-1]
		})
	} else {
		//обычный порядок
		sort.Slice(res, func(i, j int) bool {
			//проверка на существование столбца
			if len(res[i]) < column {
				return true
			} else if len(res[j]) < column {
				return false
			}

			return res[i][column-1] < res[j][column-1]
		})
	}

	return res
}

func numericSortColumn(strSlice [][]string, column int, reverse bool) [][]string {
	res := make([][]string, 0, len(strSlice))
	res = append(res, strSlice...)

	//если сортировать в обратном порядке
	if reverse {
		sort.Slice(res, func(i, j int) bool {
			//проверка на существование столбца
			if len(res[i]) < column {
				return false
			} else if len(res[j]) < column {
				return true
			}
			//преобразовываем значение в столбце к числу
			el1, err := strconv.Atoi(res[i][column-1])
			if err != nil {
				return false
			}
			el2, err := strconv.Atoi(res[j][column-1])
			if err != nil {
				return true
			}
			return el1 > el2
		})
	} else {
		//обычный порядок
		sort.Slice(res, func(i, j int) bool {
			//проверка на существование столбца
			if len(res[i]) < column {
				return true
			} else if len(res[j]) < column {
				return false
			}
			//преобразовываем значение в столбце к числу
			el1, err := strconv.Atoi(res[i][column-1])
			if err != nil {
				return true
			}
			el2, err := strconv.Atoi(res[j][column-1])
			if err != nil {
				return false
			}
			return el1 < el2

		})
	}

	return res
}

func sortUtil(strgs []string, fg *flagSettings) ([]string, error) {
	strSlice := make([]string, 0, len(strgs))
	if fg.unduplicate {
		strSlice = removeDuplicate(strgs)
	} else {
		strSlice = append(strSlice, strgs...)
	}
	res := make([]string, 0, len(strSlice))

	//если задана колонка
	if isFlagPassed("k") {
		//нумерация с 1
		if fg.column <= 0 {
			return nil, fmt.Errorf("%s: fields are numbered from 1", os.Args[0])
		}
		//разделяем на столбцы по разделителю \t
		resColumn := make([][]string, 0, len(strSlice))
		for _, el := range strSlice {
			resColumn = append(resColumn, strings.Split(el, "\t"))
		}
		//если числовая сортировка
		if fg.numeric {
			resColumn = numericSortColumn(resColumn, fg.column, fg.reverse)
		} else {
			resColumn = stringSortColumn(resColumn, fg.column, fg.reverse)
		}
		//преобразовываем обратно в []string
		for _, el := range resColumn {
			res = append(res, strings.Join(el, "\t"))
		}
	} else if fg.numeric {
		//если числовая сортировка без выбора столбца
		//разделяем на столбцы по разделителю \t
		resColumn := make([][]string, 0, len(strSlice))
		for _, el := range strSlice {
			resColumn = append(resColumn, strings.Split(el, "\t"))
		}
		resColumn = numericSortColumn(resColumn, 1, fg.reverse)
		//преобразовываем обратно в []string
		for _, el := range resColumn {
			res = append(res, strings.Join(el, "\t"))
		}
	} else {
		res = append(res, strSlice...)
		if fg.reverse {
			sort.Sort(sort.Reverse(sort.StringSlice(res)))
		} else {
			sort.Strings(res)
		}
	}

	return res, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	fg := &flagSettings{}
	flag.IntVar(&fg.column, "k", 0, "number of column for sotring")
	flag.BoolVar(&fg.numeric, "n", false, "sorting numbers")
	flag.BoolVar(&fg.reverse, "r", false, "reverse sort")
	flag.BoolVar(&fg.unduplicate, "u", false, "don't print duplicate strings")

	flag.Parse()
	strgs, err := input()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: [%s]\n", err)
		os.Exit(1)
	}

	res, err := sortUtil(strgs, fg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: [%s]\n", err)
		os.Exit(1)
	}
	for _, el := range res {
		fmt.Println(el)
	}

}
