package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		_, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
