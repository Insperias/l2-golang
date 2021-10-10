package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flagSettings struct {
	column    int
	delim     string
	separator bool
}

func input(fg *flagSettings) ([]string, error) {
	if !isFlagPassed("f") {
		return nil, fmt.Errorf("%s: option requires an argument -- 'f'", os.Args[0])

	} else if fg.column <= 0 {
		return nil, fmt.Errorf("%s: fields are numbered from 1", os.Args[0])
	}

	strgs := make([]string, 0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		strgs = append(strgs, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error: reading: err: [%s]", err)
	}

	return strgs, nil
}

func cut(fg *flagSettings, strgs []string) ([]string, error) {

	strSlice := make([][]string, 0)

	for _, el := range strgs {
		if fg.separator {
			if strings.Contains(el, fg.delim) {
				strSlice = append(strSlice, strings.Split(el, fg.delim))
			} else {
				strSlice = append(strSlice, []string{""})
			}
		} else {
			strSlice = append(strSlice, strings.Split(el, fg.delim))
		}
	}
	res := make([]string, 0, len(strSlice))
	for _, el := range strSlice {
		if fg.column > len(el) {
			res = append(res, "")
		} else {
			res = append(res, el[fg.column-1])
		}
	}

	return res, nil
}

func main() {

	var fg = &flagSettings{}
	flag.IntVar(&fg.column, "f", 0, "number of column. (Required)")
	flag.StringVar(&fg.delim, "d", "\t", "delimeter")
	flag.BoolVar(&fg.separator, "s", false, "only strings with delimeter")
	flag.Parse()

	strgs, err := input(fg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: [%s]\n", err)
		os.Exit(1)
	}

	res, err := cut(fg, strgs)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: [%s]\n", err)
		os.Exit(1)
	}

	for _, el := range res {
		fmt.Fprintln(os.Stdout, el)
	}
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
