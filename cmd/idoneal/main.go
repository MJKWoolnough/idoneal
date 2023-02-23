package main

import (
	"flag"
	"fmt"
	"os"

	"vimagination.zapto.org/idoneal/pkg/meta"
)

func main() {
	var countSequences bool

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [OPTIONS] FILE

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

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %s\n", err)
		os.Exit(2)
	}
	defer f.Close()

	if countSequences {
		count, err := meta.CountSequences(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while counting sequences: %s\n", err)
			os.Exit(2)
		}
		fmt.Fprintln(os.Stdout, count)
	}
}
