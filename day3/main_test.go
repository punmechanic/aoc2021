package main

import (
	"bufio"
	_ "embed"
	"strings"
	"testing"

	"github.com/trinitroglycerin/aoc2021/pkg/bitset"
)

//go:embed test1.txt
var inputs string

func requireSets(t *testing.T) []bitset.BitSet {
	sets, err := fromReader(bufio.NewReader(strings.NewReader(inputs)))
	if err != nil {
		t.Fatal(err)
	}

	return sets
}

func TestGammaAndEpsilon(t *testing.T) {
	sets := requireSets(t)
	gammaBits := bitset.MostCommon(sets)
	if gammaBits.Uint() != 0b10110 {
		t.Errorf("gamma was %#05b, expected 0b10110", gammaBits.Uint())
	}

	if gammaBits.Negate().Uint() != 0b01001 {
		t.Errorf("epsilon was %#05b, expected 0b01001", gammaBits.Negate().Uint())
	}
}

func TestFindOxygenGeneratorRating(t *testing.T) {
	sets := requireSets(t)
	oxy := findOxygenGeneratorRating(sets)
	if oxy != 23 {
		t.Errorf("oxy was %d, expected %d", oxy, 23)
	}
}

func TestFindCO2ScrubberRating(t *testing.T) {
	sets := requireSets(t)
	co2 := findCO2ScrubberRating(sets)
	if co2 != 10 {
		t.Errorf("co2 was %d, expected %d", co2, 10)
	}
}
