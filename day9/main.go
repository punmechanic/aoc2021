package main

import (
	"flag"
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
			if isSmallest(cell, grid.FindSiblingNodes(gr, grid.Point{X: cellIdx, Y: rowIdx})) {
				lps = append(lps, cell)
			}
		}
	}

	return lps
}

var part = flag.String("part", "", "The part of the puzzle (part-1, part-2)")

func main() {
	flag.Parse()
	gr, err := grid.Read(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	switch *part {
	case "":
		fallthrough
	case "part-1":
		sum := 0
		for _, point := range findLowPoints(gr) {
			sum += 1 + int(point)
		}

		fmt.Printf("%d\n", sum)
	case "part-2":
		basins := findLargestBasins(gr, 3)
		product := basins[0].Length
		for _, basin := range basins[1:] {
			product *= basin.Length
		}

		fmt.Printf("%d\n", product)
	default:
		log.Fatalf("invalid part %q", *part)
	}
}
