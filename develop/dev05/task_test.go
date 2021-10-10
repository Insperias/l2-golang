package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrepAfterFixed(t *testing.T) {
	fg := &flags{After: 1}
	pattern := "abc"
	reader, err := os.Open("test.txt")
	assert.Nil(t, err)

	res, err := Grep(reader, pattern, fg)
	assert.Nil(t, err)

}
