package main

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	fname := "sort_test.txt"
	expect := []string{
		"5\t55\t4",
		"5\t21\t7",
		"5\t7\t1",
	}

	res, err := readFile(fname)
	if err != nil {
		t.Errorf("error: input error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}
}
func TestDefault(t *testing.T) {
	testStr := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}
	expect := []string{
		"aa\taba\tcdd",
		"ab\tca\tdf",
		"ba\tcd\tab",
	}

	fg := &flagSettings{}
	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}
}

func TestNumeric(t *testing.T) {
	fg := &flagSettings{numeric: true}

	testStr := []string{
		"5\t23\t4",
		"32\t3\t1",
		"11\t1\t2",
	}

	expect := []string{
		"5\t23\t4",
		"11\t1\t2",
		"32\t3\t1",
	}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}

}

func TestReverse(t *testing.T) {

	testStr := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}
	expect := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}

	fg := &flagSettings{reverse: true}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}
}

func TestUnduplicate(t *testing.T) {
	testStr := []string{
		"ba\tcd\tab",
		"aa\taba\tcdd",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}
	expect := []string{
		"aa\taba\tcdd",
		"ab\tca\tdf",
		"ba\tcd\tab",
	}

	fg := &flagSettings{unduplicate: true}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}

}

func TestNumericReverse(t *testing.T) {
	fg := &flagSettings{
		reverse: true,
		numeric: true,
	}
	testStr := []string{
		"5\t23\t4",
		"32\t3\t1",
		"11\t1\t2",
	}

	expect := []string{
		"32\t3\t1",
		"11\t1\t2",
		"5\t23\t4",
	}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}

}

func TestColumn(t *testing.T) {
	os.Args = []string{"./task", "-k", "1"}
	fg := &flagSettings{}
	flag.IntVar(&fg.column, "k", 0, "number of column for sotring")

	flag.Parse()

	testStr := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}
	expect := []string{
		"aa\taba\tcdd",
		"ab\tca\tdf",
		"ba\tcd\tab",
	}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}
}

func TestColumnReverse(t *testing.T) {
	os.Args = []string{"./task", "-k", "1"}
	fg := &flagSettings{
		column:  1,
		reverse: true,
	}
	flag.Parse()

	testStr := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}
	expect := []string{
		"ba\tcd\tab",
		"ab\tca\tdf",
		"aa\taba\tcdd",
	}

	res, err := sortUtil(testStr, fg)
	if err != nil {
		t.Errorf("error: sort error [%s]", err)
	}

	if strings.Join(expect, " ") != strings.Join(res, " ") {
		t.Errorf("error: expect [%s], got [%s]", expect, res)
	}
}
