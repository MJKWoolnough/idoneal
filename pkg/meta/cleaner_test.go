package meta

import (
	"io"
	"strings"
	"testing"
)

func TestCleaner(t *testing.T) {
	var buf [5]byte
	for n, test := range [...]struct {
		Input, Output string
	}{
		{
			"A",
			"A",
		},
		{
			"\n",
			"",
		},
		{
			"A\nB",
			"A\nB",
		},
		{
			"\n\nA\nB\n\n",
			"A\nB\n",
		},
		{
			"A\n\nB",
			"A\nB",
		},
		{
			"A\r\nB",
			"A\nB",
		},
		{
			"A\r\n\nB",
			"A\nB",
		},
		{
			"A\r\n\r\n\r\n\n\n\nB\n\n\nC\r\n\n\r\nD",
			"A\nB\nC\nD",
		},
	} {
		for i := 1; i <= 5; i++ {
			var sb strings.Builder
			io.CopyBuffer(&sb, cleaner(strings.NewReader(test.Input)), buf[:i])
			if output := sb.String(); test.Output != output {
				t.Errorf("test %d.%d: expecting output %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
