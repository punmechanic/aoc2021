package main

import (
	"bytes"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
	"github.com/trinitroglycerin/aoc2021/pkg/grid"
)

//go:embed test1.txt
var testData []byte

func TestEnumerateLowPoints(t *testing.T) {
	gr, err := grid.Read(bytes.NewBuffer(testData))
	assert.NoError(t, err)

	assert.Equal(t, []uint8{1, 0, 5, 5}, findLowPoints(gr))
}
