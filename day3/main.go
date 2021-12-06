package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"

	"github.com/trinitroglycerin/aoc2021/pkg/bitset"
)

type lineReader interface {
	ReadLine() ([]byte, bool, error)
}

func fromReader(r lineReader) ([]bitset.BitSet, error) {
	var bitSets []bitset.BitSet
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		bitSets = append(bitSets, bitset.FromString(string(line)))
	}

	return bitSets, nil
}

func findOxygenGeneratorRating(sets []bitset.BitSet) (uint, error) {
	// Determine the most common value in the current bit posiion, and keep only numbers with that bit in that position.
	// If 0 and 1 are equally common, keep values with a 1 in the position being considered.
	return 0, nil
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var bitSets []bitset.BitSet
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		bitSets = append(bitSets, bitset.FromString(string(line)))
	}

	gammaBits := bitset.MostCommon(bitSets)
	println(gammaBits.Uint() * gammaBits.Negate().Uint())
}
