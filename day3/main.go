package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"

	"github.com/trinitroglycerin/aoc2021/pkg/bitset"
)

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
