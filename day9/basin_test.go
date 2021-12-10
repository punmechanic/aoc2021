package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasinListSortsDescending(t *testing.T) {
	basins := basinList{
		{Length: 5},
		{Length: 7},
		{Length: 3},
		{Length: 9},
	}

	sort.Sort(basins)
	assert.Equal(t, basin{Length: 9}, basins[0])
	assert.Equal(t, basin{Length: 7}, basins[1])
	assert.Equal(t, basin{Length: 5}, basins[2])
	assert.Equal(t, basin{Length: 3}, basins[3])
}
