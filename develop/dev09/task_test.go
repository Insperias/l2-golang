package main

import (
	"os"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	url := "https://www.wildberries.ru"

	err := wget(url, "")
	if err != nil {
		t.Errorf("can't download: %s", err)
	}
}

func TestDownloadFile(t *testing.T) {
	url := "https://golang.org/dl/go1.15.2.linux-amd64.tar.gz"

	err := wget(url, "")
	if err != nil {
		t.Errorf("can't download: %s", err)
	}
}

func TestOutputFile(t *testing.T) {
	url := "https://www.wildberries.ru"
	outputFile := "wildberries.html"

	err := wget(url, outputFile)
	if err != nil {
		t.Errorf("can't download: %s", err)
	}

	_, err = os.Stat(outputFile)
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("expect saving in %s file : %s", outputFile, err)
		}
	}
}

func TestResolveIdenticalFilenames(t *testing.T) {
	url := "https://www.wildberries.ru"

	i := 0
	for i < 20 {
		err := wget(url, "")
		if err != nil {
			t.Errorf("can't download: %s", err)
		}
		i++
	}
}
