package meta

import (
	"strings"
	"testing"
)

func TestCountNucleotides(t *testing.T) {
	for n, test := range [...]struct {
		Input string
		Count int
	}{
		{
			"",
			0,
		},
		{
			"@\nAAA",
			3,
		},
		{
			"@\nABC\n+\n!!!\n@\nDEFG",
			7,
		},
		{
			"@\nABC\n+\n!!!\n@\nDEFG\n+\n!!!!\n@\n987654321\n+\n",
			16,
		},
	} {
		if count, err := CountNucleotides(strings.NewReader(test.Input)); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else if count != test.Count {
			t.Errorf("test %d: expecting to count %d nucleotides, got %d", n+1, test.Count, count)
		}
	}
}
