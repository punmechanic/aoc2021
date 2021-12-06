package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	first := true
	var acc uint
	var prev uint16
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		raw, err := strconv.ParseUint(string(line), 10, 16)
		if err != nil {
			log.Fatalf("failed to parse %q: %s", line, err)
		}

		current := uint16(raw)
		if current > prev && !first {
			acc++
		}

		prev = current
		first = false
	}

	println(acc)
}
