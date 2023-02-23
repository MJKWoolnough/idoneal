package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"vimagination.zapto.org/idoneal/tests"
)

var exitCode int

func init() {
	Exit = func(code int) {
		exitCode = code
	}
	FS = tests.FS
	flagErrorHandler = flag.ContinueOnError
}

func testOutput(t *testing.T, flag, ext string) {
	t.Helper()
	for n, test := range [...]string{"a", "b", "c", "d", "e"} {
		var stdout, stderr strings.Builder
		Stdout = &stdout
		Stderr = &stderr
		os.Args = append(os.Args[:1], flag, fmt.Sprintf("%s.fastq", test))
		main()
		if stderr.Len() > 0 {
			t.Errorf("test %d: unexpected error: %s", n+1, stderr.String())
		} else {
			f, _ := FS.Open(fmt.Sprintf("%s.%s", test, ext))
			expected, _ := io.ReadAll(f)
			if output := stdout.String(); output != string(expected) {
				t.Errorf("test %d: expecting output %q, got %q", n+1, expected, output)
			}
		}
	}
}

func TestMainCountSequences(t *testing.T) {
	testOutput(t, "-s", "sequences")
}

func TestMainCountSequencesLong(t *testing.T) {
	testOutput(t, "--sequences", "sequences")
}

func TestMainCountNucleotides(t *testing.T) {
	testOutput(t, "-n", "nucleotides")
}

func TestMainCountNucleotidesLong(t *testing.T) {
	testOutput(t, "--nucleotides", "nucleotides")
}
