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

type Board [5][5]uint

func visitLines(r *bufio.Reader, visitor func(line []byte) error) error {
	for {
		line, _, err := r.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		if err := visitor(line); err != nil {
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

func (st *bingoGameReader) readSequence(line []byte) error {
	for _, substr := range strings.Split(string(line), ",") {
		digit, err := strconv.ParseUint(substr, 10, 16)
		if err != nil {
			return err
		}

		st.Sequence = append(st.Sequence, uint(digit))
	}

	return nil
}

func (st *bingoGameReader) Read(r io.Reader) error {
	br := bufio.NewReader(r)
	return visitLines(br, func(line []byte) error {
		if !st.sequenceRead {
			if err := st.readSequence(line); err != nil {
				return err
			}
			st.sequenceRead = true
		}

		return nil
	})
}

func main() {
	var game bingoGameReader
	if err := game.Read(os.Stdin); err != nil {
		log.Fatal(err)
	}
}
