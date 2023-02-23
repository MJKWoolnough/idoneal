package meta

import (
	"fmt"
	"io"
)

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
