package main

import (
	"fmt"
	"log"
	"os"

	"github.com/trinitroglycerin/aoc2021/pkg/grid"
)

func isSmallest(a grid.Node, bs []grid.Node) bool {
	for _, b := range bs {
		if a >= b {
			return false
		}
	}

	return true
}

func findLowPoints(gr grid.Grid) []grid.Node {
	var lps []grid.Node
	for rowIdx, row := range gr {
		for cellIdx, cell := range row {
			if isSmallest(cell, grid.FindSiblingNodes(gr, rowIdx, cellIdx)) {
				lps = append(lps, cell)
			}
		}
	}

	return lps
}

func main() {
	gr, err := grid.Read(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, point := range findLowPoints(gr) {
		sum += 1 + int(point)
	}

	fmt.Printf("%d\n", sum)
}
