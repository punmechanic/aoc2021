package main

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed test1.txt
var inputs string

func findOxygenGeneratorRating(r io.Reader) (uint, error) {
	return 0, nil
}

func findCO2ScrubberRating(r io.Reader) (uint, error) {
	return 0, nil
}

func TestFindOxygenGeneratorRating(t *testing.T) {
	r := strings.NewReader(inputs)
	oxy, err := findOxygenGeneratorRating(r)
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
