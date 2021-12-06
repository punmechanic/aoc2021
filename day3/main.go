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
	// Determine the least common value in the current bit position, and keep only numbers with that bit in that position.
	// If 0 and 1 are equally common, keep values with a 0 in the position being considered.
	// This is similar to bitset.MostCommon but we gradually restrict the search list.
	for i := 0; i < sets[0].BitLength(); i++ {
		var ons, offs []bitset.BitSet

		if len(sets) == 1 {
			break
		}

		for _, set := range sets {
			if set[i] {
				ons = append(ons, set)
			} else {
				offs = append(offs, set)
			}
		}

		if len(ons) >= len(offs) {
			sets = offs
		} else {
			sets = ons
		}
	}

	return sets[0].Uint()
}

func findOxygenGeneratorRating(sets []bitset.BitSet) uint {
	// Determine the most common value in the current bit position, and keep only numbers with that bit in that position.
	// If 0 and 1 are equally common, keep values with a 1 in the position being considered.
	// This is similar to bitset.MostCommon but we gradually restrict the search list.
	for i := 0; i < sets[0].BitLength(); i++ {
		var ons, offs []bitset.BitSet
		if len(sets) == 1 {
			break
		}

		for _, set := range sets {
			if set[i] {
				ons = append(ons, set)
			} else {
				offs = append(offs, set)
			}
		}

		if len(ons) >= len(offs) {
			sets = ons
		} else {
			sets = offs
		}
	}

	return sets[0].Uint()
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
