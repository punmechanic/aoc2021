package grid

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Node = uint8

type Grid = [][]Node

func Read(r io.Reader) (Grid, error) {
	var grid Grid
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		var row []Node
		for _, b := range line {
			if b == '\n' {
				break
			}

			n, err := strconv.ParseUint(string(b), 10, 8)
			if err != nil {
				return nil, err
			}

			row = append(row, uint8(n))
		}

		grid = append(grid, row)
	}

	return grid, nil
}

func Get(grid Grid, rowIdx, cellIdx int) (Node, bool) {
	if rowIdx > len(grid)-1 {
		return 0, false
	}

	slice := grid[rowIdx]
	if cellIdx > len(slice)-1 {
		return 0, false
	}

	return slice[cellIdx], true
}

func FindSiblingNodes(grid Grid, rowIdx, cellIdx int) []Node {
	var siblings []Node
	places := [][]int{
		{rowIdx - 1, cellIdx},
		{rowIdx + 1, cellIdx},
		{rowIdx, cellIdx + 1},
		{rowIdx, cellIdx - 1},
	}

	for _, place := range places {
		rowIdx := place[0]
		cellIdx := place[1]
		if cellIdx < 0 || rowIdx < 0 {
			continue
		}

		if node, ok := Get(grid, rowIdx, cellIdx); ok {
			siblings = append(siblings, node)
		}
	}

	return siblings
}
