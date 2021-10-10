package main

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestStandartCut(t *testing.T) {
	testStr := []string{
		"sek\tlek",
		"brak\tmak",
		"mir\trim",
	}
	expect := []string{"sek", "brak", "mir"}

	fg := &flagSettings{
		column:    1,
		delim:     "\t",
		separator: false,
	}

	res, _ := cut(fg, testStr)
	if strings.Join(res, " ") != strings.Join(expect, " ") {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}

}

func TestFlagF(t *testing.T) {

	fg := &flagSettings{}

	_, err := input(fg)
	if err == nil {
		t.Errorf("invalid result: expect error, got nil")
	}
}

func TestNegativeColumn(t *testing.T) {
	os.Args = []string{"./task", "-f", "-1"}
	fg := &flagSettings{}
	flag.IntVar(&fg.column, "f", 0, "number of column. (Required)")

	flag.Parse()
	_, err := input(fg)
	if err == nil {
		t.Errorf("invalid result: expect error, got nil [%s]", err)
	}

}

func TestColumnOutRange(t *testing.T) {
	testStr := []string{
		"sek\tlek\tbek",
		"brak\tmak",
		"mir\trim",
	}
	expect := []string{"bek", "", ""}

	fg := &flagSettings{
		column:    3,
		delim:     "\t",
		separator: false,
	}

	res, _ := cut(fg, testStr)
	if strings.Join(res, " ") != strings.Join(expect, " ") {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestDelimeter(t *testing.T) {
	fg := &flagSettings{
		column:    1,
		delim:     " ",
		separator: false,
	}
	testStr := []string{
		"sek lek",
		"brak mak",
		"mir rim",
	}
	expect := []string{"sek", "brak", "mir"}
	res, _ := cut(fg, testStr)
	if strings.Join(res, " ") != strings.Join(expect, " ") {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestSeparator(t *testing.T) {
	fg := &flagSettings{
		column:    1,
		delim:     "\t",
		separator: true,
	}

	testStr := []string{
		"sek\tlek",
		"brak mak",
		"mir\trim",
	}
	expect := []string{"sek", "", "mir"}

	res, _ := cut(fg, testStr)
	if strings.Join(res, " ") != strings.Join(expect, " ") {
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}

}
