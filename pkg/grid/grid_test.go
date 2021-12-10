package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGrid Grid = [][]uint8{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func sequentialDecrease(prev, next Node) bool {
	return next-1 == prev
}

func TestSpiderNodesCanIdentifyBasins(t *testing.T) {
	assert.Len(t, SpiderNodes(testGrid, Point{X: 1, Y: 0}, sequentialDecrease), 3)
	assert.Len(t, SpiderNodes(testGrid, Point{X: 9, Y: 0}, sequentialDecrease), 9)
	assert.Len(t, SpiderNodes(testGrid, Point{X: 2, Y: 2}, sequentialDecrease), 14)
	assert.Len(t, SpiderNodes(testGrid, Point{X: 6, Y: 4}, sequentialDecrease), 9)
}

// nodes is a package level variable used to prevent the compiler optimising away the results of benchmarks
var nodes []Node

func BenchmarkSpiderNodesLargestSet(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		nodes = SpiderNodes(testGrid, Point{X: 2, Y: 2}, sequentialDecrease)
	}
}

func BenchmarkSpiderNodesSmallestSet(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		nodes = SpiderNodes(testGrid, Point{X: 1, Y: 0}, sequentialDecrease)
	}
}
