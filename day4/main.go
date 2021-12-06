package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type Board [5][5]uint

func visitLines(r *bufio.Reader, visitor func(line []byte)) error {
	for {
		_, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

type bingoGameReader struct {
	sequenceRead bool
	Sequence     []uint
	Boards       []Board
}

func (st *bingoGameReader) Read(r io.Reader) error {
	br := bufio.NewReader(r)

	return visitLines(br, func(line []byte) {
		// TODO: Load game here
	})
}

func main() {
	var game bingoGameReader
	if err := game.Read(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
