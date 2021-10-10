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

func NewRegexChecker(r *regexp.Regexp) *regexChecker {
	return &regexChecker{R: r}
}

func (c *regexChecker) check(s string) bool {
	return c.R.MatchString(s)
}

type equalChecker struct {
	pat string
}

func NewEqualChecker(pat string) *equalChecker {
	return &equalChecker{pat: pat}
}

func (c *equalChecker) check(s string) bool {
	return strings.Contains(s, c.pat)
}

func Grep(r io.Reader, pattern string, fg *flags) (string, error) {
	var chr checker

	if fg.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	if fg.Fixed {
		chr = NewEqualChecker(pattern)
	} else {
		rx, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		chr = NewRegexChecker(rx)
	}

	lt := fg.Context
	rt := fg.Context
	if fg.Before > 0 {
		lt = fg.Before
	}
	if fg.After > 0 {
		rt = fg.After
	}

	strgs := make(map[int]bool)
	var resIds []int
	var allStr []string

	var counter, i int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		txt := sc.Text()
		allStr = append(allStr, txt)
		if fg.IgnoreCase {
			txt = strings.ToLower(txt)
		}

		if chr.check(txt) {
			counter++
			if lt > 0 {
				for j := i - lt; j < i; j++ {
					if _, ok := strgs[j]; !ok {
						resIds = append(resIds, j)
						strgs[j] = false
					}
				}
			}
			if rt > 0 {
				for j := i + rt; j > i; j-- {
					if _, ok := strgs[j]; !ok {
						resIds = append(resIds, j)
						strgs[j] = false
					}
				}
			}

			if _, ok := strgs[i]; !ok {
				resIds = append(resIds, i)
			}
			strgs[i] = true
		}
		i++

	}

	if fg.Count {
		return fmt.Sprintf("%v matches", counter), nil
	}

	sort.Ints(resIds)

	var res []string
	lIndex := ""
	if !fg.Invert {
		match := "  "
		for _, k := range resIds {
			if k > -1 && k < len(allStr) {
				if strgs[k] {
					match = "* "
				} else {
					match = " "
				}
				if fg.LineNumber {
					lIndex = strconv.Itoa(k) + " "
				}
				res = append(res, fmt.Sprintf("%v%v%v", lIndex, match, allStr[k]))
			}
		}
	} else {
		for i, v := range allStr {
			if _, ok := strgs[i]; !ok {
				if fg.LineNumber {
					lIndex = strconv.Itoa(i) + " "
				}
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

	reader, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("grep err: ", err.Error())
		return
	}

	pattern := flag.Arg(0)

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		fmt.Println("grep err: ", err.Error())
		return
	}

	fmt.Println(res)
}
