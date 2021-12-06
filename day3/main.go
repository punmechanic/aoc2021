package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

const (
	ONE uint8 = 49
)

func reconstituteUint(bits []bool) uint {
	var o uint
	shift := len(bits) - 1
	for _, bit := range bits {
		var b uint8
		if bit == true {
			b = 1
		}

		shifted := uint(b) << shift
		o |= shifted
		shift--
	}

	return o
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var lines uint
	// Rather than attempt to calculate the epsilon and gamma at the same time, we can count the number of 1s we read for each bit.
	// If the number of 1s is larger than at least half the number of lines, then we know that 1 is the most common bit, and 0 the least.
	var oneCounts []uint
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if lines == 0 {
			oneCounts = make([]uint, len(line))
		}

		for idx, char := range line {
			if char == ONE {
				oneCounts[idx]++
			}
		}

		lines++
	}

	half := lines / 2
	gammaBits := make([]bool, len(oneCounts))
	epsilonBits := make([]bool, len(oneCounts))
	for idx, count := range oneCounts {
		if count > half {
			gammaBits[idx] = true
		} else {
			epsilonBits[idx] = true
		}
	}

	gamma := reconstituteUint(gammaBits)
	epsilon := reconstituteUint(epsilonBits)
	println(gamma * epsilon)
}
