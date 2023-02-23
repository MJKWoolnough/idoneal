// This package contains a CLI program to extract basic metadata from FASTQ formatted files.
package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"vimagination.zapto.org/idoneal/pkg/meta"
)

type OS struct{}

func (OS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

// Vars to be mocked for testing
var (
	Stdout           io.Writer = os.Stdout
	Stderr           io.Writer = os.Stderr
	FS               fs.FS     = OS{}
	Exit                       = os.Exit
	flagErrorHandler           = flag.ExitOnError
)

func openFile(file string) (io.ReadCloser, error) {
	var (
		f   io.ReadCloser
		err error
	)
	f, err = FS.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w\n", err)
	}
	if strings.HasSuffix(file, ".gz") {
		f, err = gzip.NewReader(f)
		if err != nil {
			return nil, fmt.Errorf("error starting decompression: %w", err)
		}
	}
	return f, nil
}

func main() {
	var countSequences, countNucleotides bool

	flags := flag.NewFlagSet(os.Args[0], flagErrorHandler)
	flags.SetOutput(Stderr)
	flags.Usage = func() {
		fmt.Fprintf(Stderr, `Usage: %s [OPTIONS] FILE

OPTIONS:
  -h, --help          Print this help.
  -n, --nucleotides   Print the number of nucleotides in the FILE.
  -s, --sequences     Print the number of sequences in the FILE.
`, os.Args[0])
	}
	flags.BoolVar(&countSequences, "s", false, "")
	flags.BoolVar(&countSequences, "sequences", false, "")
	flags.BoolVar(&countNucleotides, "n", false, "")
	flags.BoolVar(&countNucleotides, "nucleotides", false, "")
	flags.Parse(os.Args[1:])

	file := flags.Arg(0)

	if file == "" {
		flags.Usage()
		return
	}

	if countSequences {
		f, err := openFile(file)
		if err != nil {
			fmt.Fprintln(Stderr, err)
		}
		defer f.Close()
		count, err := meta.CountSequences(f)
		if err != nil {
			fmt.Fprintf(Stderr, "error while counting sequences: %s\n", err)
			Exit(2)
			return
		}
		fmt.Fprintln(Stdout, count)
	}
	if countNucleotides {
		f, err := openFile(file)
		if err != nil {
			fmt.Fprintln(Stderr, err)
		}
		defer f.Close()
		count, err := meta.CountNucleotides(f)
		if err != nil {
			fmt.Fprintf(Stderr, "error while counting nucleotides: %s\n", err)
			Exit(2)
			return
		}
		fmt.Fprintln(Stdout, count)
	}
}
