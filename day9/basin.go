package main

import (
	"sort"

	"github.com/trinitroglycerin/aoc2021/pkg/grid"
)

type basin struct {
	Length int
}

type basinList []basin

func (l basinList) Len() int {
	return len(l)
}

func (l basinList) Less(i, j int) bool {
	return l[i].Length > l[j].Length
}

func (l basinList) Swap(i, j int) {
	a := l[i]
	l[i] = l[j]
	l[j] = a
}

func findBasins(gr grid.Grid) basinList {
	var basins basinList
	sequentialDecrease := func(prev, next grid.Node) bool {
		return next-1 == prev
	}

	for _, root := range grid.Spider(gr, sequentialDecrease) {
		basins = append(basins, basin{Length: len(root)})
	}

	return basins
}

func findLargestBasins(gr grid.Grid, n int) []basin {
	basins := findBasins(gr)
	sort.Sort(basins)
	return basins[:n]
}
