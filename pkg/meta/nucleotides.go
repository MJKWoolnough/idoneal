package meta

import (
	"fmt"
	"io"
)

// CountNucleotides returns the number of Nucleotides in the passed FASTQ
// formatted Reader.
//
// NB: No validation is performed on the data being read.
func CountNucleotides(r io.Reader) (int, error) {
	lr := cleaner(r)
	var buf [1024]byte
	line := 0
	count := 0
	for {
		n, err := lr.Read(buf[:])
		for _, b := range buf[:n] {
			if b == '\n' {
				line = (line + 1) % 4 // what sequence line are we on?
			} else if line == 1 { // second line
				count++
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, fmt.Errorf("error while counting nucleotides: %w", err)
		}
	}
	return count, nil
}
