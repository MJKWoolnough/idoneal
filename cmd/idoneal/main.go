package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"

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

func main() {
	var countSequences bool

	flags := flag.NewFlagSet(os.Args[0], flagErrorHandler)
	flags.SetOutput(Stderr)
	flags.Usage = func() {
		fmt.Fprintf(Stderr, `Usage: %s [OPTIONS] FILE

OPTIONS:
  -h, --help       Print this help.
  -s, --sequences  Print the number of sequences in the FILE.
`, os.Args[0])
	}
	flags.BoolVar(&countSequences, "s", false, "")
	flags.BoolVar(&countSequences, "sequences", false, "")
	flags.Parse(os.Args[1:])

	file := flags.Arg(0)

	if file == "" {
		flags.Usage()
		return
	}

	f, err := FS.Open(file)
	if err != nil {
		fmt.Fprintf(Stderr, "error opening file: %s\n", err)
		Exit(2)
		return
	}
	defer f.Close()

	if countSequences {
		count, err := meta.CountSequences(f)
		if err != nil {
			fmt.Fprintf(Stderr, "error while counting sequences: %s\n", err)
			Exit(2)
			return
		}
		fmt.Fprintln(Stdout, count)
	}
}