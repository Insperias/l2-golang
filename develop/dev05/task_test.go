package main

import (
	"os"
	"testing"
)

func TestAfterFixed(t *testing.T) {
	fg := &flags{
		After: 1,
		Fixed: true,
	}
	pattern := "asd(a|b|c)"
	expect := "+ asd(a|b|c)\n qwertaabc"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect{
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestBeforeIgnoreCase(t *testing.T){
fg := &flags{
	Before: 1,
	IgnoreCase: true,
}
	pattern := "asd(a|b|c)"
	expect := " AsD(A|b|C)\n+ acdasdasc\n abbaaaasf\n+ acdaSdasc\n+ acdaSdcasc"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect{
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}

}

func TestCount(t *testing.T){
	fg := &flags{Count: true}
	pattern := "abc"
	expect := "2 matches"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect{
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestLineNumberContext(t *testing.T){
	fg := &flags{
		Context: 1,
		LineNumber: true,
	}
	pattern := "ks"
	expect := "7  acdaSdcasc\n8 + artksddas\n9  asd(a|b|c)"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect{
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}

func TestInvertIgnoreCase(t *testing.T){
	fg := &flags{
		Invert: true,
		IgnoreCase: true,
	}
	pattern := "asd"
	expect := "abcabcabc\naaabbbccc\nabbaaaasf\nartksddas\nqwertaabc"
	reader, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("open file error: [%s]", err)
	}

	res, err := Grep(reader, pattern, fg)
	if err != nil {
		t.Errorf("grep error: [%s]", err)
	}

	if res != expect{
		t.Errorf("invalid result: expect [%s], got [%s]", expect, res)
	}
}
