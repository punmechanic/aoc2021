package grid

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Node = uint8

type Grid = [][]Node

type Point struct {
	X, Y int
}

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

func Get(grid Grid, pt Point) (Node, bool) {
	if pt.Y < 0 || pt.X < 0 {
		return 0, false
	}

	if pt.Y > len(grid)-1 {
		return 0, false
	}

	slice := grid[pt.Y]
	if pt.X > len(slice)-1 {
		return 0, false
	}

	return slice[pt.X], true
}

func FindSiblingNodes(grid Grid, pt Point) []Node {
	var siblings []Node
	pts := []Point{
		{X: pt.X, Y: pt.Y - 1},
		{X: pt.X, Y: pt.Y + 1},
		{X: pt.X - 1, Y: pt.Y},
		{X: pt.X + 1, Y: pt.Y},
	}

	for _, point := range pts {
		if point.X < 0 || point.Y < 0 {
			continue
		}

		if node, ok := Get(grid, point); ok {
			siblings = append(siblings, node)
		}
	}

	return siblings
}

// Spider finds siblings for _all_ nodes according to the predicate.
func Spider(grid Grid, predicate func(prev, current Node) bool) [][]Node {
	var ns [][]Node
	for y, row := range grid {
		for x := range row {
			pt := Point{X: x, Y: y}
			ns = append(ns, SpiderNodes(grid, pt, predicate))
		}
	}
	return ns
}

// SpiderNodes functions similarly to FindSiblingNodes, but continues searching in each given direction until predicate returns false.
func SpiderNodes(grid Grid, pt Point, predicate func(prev, current Node) bool) []Node {
	worker := spiderWorker{
		grid:      grid,
		checked:   make(map[Point]struct{}),
		predicate: predicate,
	}

	return worker.Spider(nil, pt)
}

type spiderWorker struct {
	grid      Grid
	checked   map[Point]struct{}
	predicate func(prev, current Node) bool
}

func (w spiderWorker) Spider(initial *Node, pt Point) []Node {
	if _, ok := w.checked[pt]; ok {
		return nil
	}

	current, ok := Get(w.grid, pt)
	if !ok {
		return nil
	}

	if current == 9 {
		return nil
	}

	if initial != nil && !w.predicate(*initial, current) {
		return nil
	}

	nodes := []Node{current}
	nodes = append(nodes, w.Spider(&current, Point{X: pt.X - 1, Y: pt.Y})...)
	nodes = append(nodes, w.Spider(&current, Point{X: pt.X + 1, Y: pt.Y})...)
	nodes = append(nodes, w.Spider(&current, Point{X: pt.X, Y: pt.Y - 1})...)
	nodes = append(nodes, w.Spider(&current, Point{X: pt.X, Y: pt.Y + 1})...)

	if len(nodes) > 1 {
		// If we found at least one other Node, then we have found the 'basin' that this Node is a part of.
		w.checked[pt] = struct{}{}
		// However, if we found zero Nodes, that means that we have not yet found the basin for this Node.
		// All Nodes must be part of at least ONE basin.
	}

	return nodes
}
