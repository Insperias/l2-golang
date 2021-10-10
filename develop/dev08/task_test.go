package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCdAndPwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dd, err := parseCommand("", "pwd")
	assert.Nil(t, err)
	assert.Equal(t, dir, dd)
	_, err = parseCommand("", "cd dir_for_test")
	assert.Nil(t, err)
	dd, err = parseCommand("", "pwd")
	assert.Nil(t, err)
	assert.Equal(t, dir+"/dir_for_test", dd)
}

func TestEcho(t *testing.T) {
	r, err := parseCommand("", "echo 123")
	assert.Nil(t, err)
	assert.Equal(t, "123", r)
}

func TestExec(t *testing.T) {
	r, err := parseCommand("", "exec echo 123")
	assert.Nil(t, err)
	assert.Equal(t, "123\n", r)
}

func TestPsAndFork(t *testing.T) {
	pid, err := parseCommand("", "fork()")
	assert.Nil(t, err)
	ps, err := parseCommand("", "ps")
	assert.Nil(t, err)
	assert.Contains(t, ps, pid)
}

func TestPipeline(t *testing.T) {
	dir1, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ff := strings.Split(dir1, "/")
	dir1 = strings.Join(ff[:len(ff)-1], "/")
	buf := bytes.NewBuffer([]byte{})
	parsePipeline(buf, "echo .. | cd")
	assert.Nil(t, err)
	dir2, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, dir2, dir1)
}
