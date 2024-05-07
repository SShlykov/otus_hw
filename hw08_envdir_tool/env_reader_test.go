package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadDir(t *testing.T) {
	path := "testdata/env"

	fileNames, err := ReadDir(path)
	fmt.Printf("%#v\n", fileNames)

	assert.Empty(t, err)
}
