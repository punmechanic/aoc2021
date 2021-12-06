package main

import (
	"bufio"
	"errors"
	"fmt"
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

func findCO2ScrubberRating(sets []bitset.BitSet) uint {
	return bitset.UniformZeroes(sets).Uint()
}

func findOxygenGeneratorRating(sets []bitset.BitSet) uint {
	return bitset.Uniform(sets).Uint()
}

func findLifeSupportRating(sets []bitset.BitSet) uint {
	return findCO2ScrubberRating(sets) * findOxygenGeneratorRating(sets)
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
	powerConsumption := gammaBits.Uint() * gammaBits.Negate().Uint()
	fmt.Printf("power consumption: %d (%#08b)\n", powerConsumption, powerConsumption)
	lifeSupport := findLifeSupportRating(bitSets)
	fmt.Printf("life support rating: %d (%#08b)\n", lifeSupport, lifeSupport)
}
