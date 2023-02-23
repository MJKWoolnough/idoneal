package meta

import (
	"strings"
	"testing"
)

func TestCountSequences(t *testing.T) {
	for n, test := range [...]struct {
		Input string
		Count int
	}{
		{
			"",
			0,
		},
		{
			"@",
			0,
		},
		{
			"@\nA",
			0,
		},
		{
			"@\nA\n+",
			0,
		},
		{
			"@\nA\n+\n!",
			1,
		},
		{
			"@\nAAAAAAAAAAAAAAAAAAAA\n+",
			0,
		},
		{
			"@\nAAAAAAAAAAAAAAAAAAAA\n+\n!!!!!!!!!!!!!!!!!!!!",
			1,
		},
		{
			"@\nA\n+\n!\n",
			1,
		},
		{
			"@\nA\n+\n!\n",
			1,
		},
		{
			"@\nA\n+\n!\n@\nB\n+",
			1,
		},
		{
			"@\nA\n+\n!\n@\nB\n+\n!",
			2,
		},
		{
			"@\nA\n+\n!\n@\nB\n+\n!\n@\nC\n+\n!\n@\nD\n+\n!\n@\nE\n+\n!\n@\nF\n+\n!\n@\nG\n+\n!\n@\nH\n+\n!\n@\nI\n+\n!\n@\nJ\n+\n!\n@\nK\n+\n!\n@\nL\n+\n!\n",
			12,
		},
	} {
		if count, err := CountSequences(strings.NewReader(test.Input)); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else if count != test.Count {
			t.Errorf("test %d: expecting to count %d sequences, got %d", n+1, test.Count, count)
		}
	}
}
