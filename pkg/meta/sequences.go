// The meta package provides some simple convenience functions for dealing with
// FASTQ files.
package meta

import (
	"fmt"
	"io"
)

// CountSequences returns the number of Sequences in the passed FASTQ formatted
// Reader.
//
// NB: This function acts naively, simply counting the non-blank lines in the
// Reader, and diving by 4.
func CountSequences(r io.Reader) (int, error) {
	lr := cleaner(r)
	var buf [1024]byte
	lineCount := 1
	for {
		n, err := lr.Read(buf[:])
		for _, b := range buf[:n] {
			if b == '\n' {
				lineCount++
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, fmt.Errorf("error while counting sequences: %w", err)
		}
	}
	return lineCount >> 2, nil
}
