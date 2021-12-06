package main

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/trinitroglycerin/aoc2021/pkg/bitset"
)

//go:embed test1.txt
var inputs string

func findCO2ScrubberRating(r io.Reader) (uint, error) {
	return 0, nil
}

func TestGammaAndEpsilon(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(inputs))
	sets, err := fromReader(r)
	if err != nil {
		t.Errorf("failed to read BitSets: %s", err)
	}

	gammaBits := bitset.MostCommon(sets)

	if gammaBits.Uint() != 0b10110 {
		t.Errorf("gamma was %#05b, expected 0b10110", gammaBits.Uint())
	}

	if gammaBits.Negate().Uint() != 0b01001 {
		t.Errorf("epsilon was %#05b, expected 0b01001", gammaBits.Negate().Uint())
	}
}

func TestFindOxygenGeneratorRating(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(inputs))
	sets, err := fromReader(r)
	if err != nil {
		t.Errorf("failed to read BitSets: %s", err)
	}

	oxy, err := findOxygenGeneratorRating(sets)
	if err != nil {
		t.Errorf("failed to find oxygen generator rating: %s", err)
	}

	if oxy != 23 {
		t.Errorf("oxy was %d, expected %d", oxy, 23)
	}
}

func TestFindCO2ScrubberRating(t *testing.T) {
	r := strings.NewReader(inputs)
	co2, err := findCO2ScrubberRating(r)
	if err != nil {
		t.Errorf("failed to find CO2 scrubber rating: %s", err)
	}

	if co2 != 10 {
		t.Errorf("co2 was %d, expected %d", co2, 10)
	}
}
