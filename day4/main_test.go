package main

import (
	_ "embed"

	"testing"
)

//go:embed test1.txt
var inputs []byte

type Board [5][5]uint

func TestCanLoadGame(t *testing.T) {
	var sequence []uint
	var boards []Board

	expectedSeq := []uint{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	expectedBoards := []Board{
		[5][5]uint{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		[5][5]uint{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		[5][5]uint{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}

	if len(sequence) != len(expectedSeq) {
		t.Errorf("incorrect sequence length - expected %d, got %d", len(expectedSeq), len(sequence))
	}

	for idx, n := range sequence {
		t.Errorf("incorrect sequence number at sequence[%d] - expected %d, got %d", idx, expectedSeq[idx], n)
	}

	if len(boards) != len(expectedBoards) {
		t.Errorf("incorrect number of boards - expected %d, got %d", len(expectedBoards), len(boards))
	}

	for boardIdx, board := range boards {
		expectedBoard := expectedBoards[boardIdx]
		for rowIdx, row := range board {
			expectedRow := expectedBoard[rowIdx]
			for cellIdx, cell := range row {
				expectedCell := expectedRow[cellIdx]
				if len(boards) != len(expectedBoards) {
					t.Errorf("incorrect cell value at boards[%d][%d][%d] - expected %d, got %d", boardIdx, rowIdx, cellIdx, expectedCell, cell)
				}
			}
		}
	}
}
