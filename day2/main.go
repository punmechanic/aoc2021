package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var depth, hpos int64
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(string(line), " ")
		v, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			log.Fatalf("failed to parse %q: %s", parts[1], err)
		}

		switch parts[0] {
		case "forward":
			hpos += v
		case "down":
			depth += v
		case "up":
			depth -= v
		default:
			log.Fatalf("unrecognised direction %q\n", parts[0])
		}
	}

	println(depth * hpos)
}
