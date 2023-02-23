# meta
--
    import "vimagination.zapto.org/idoneal/pkg/meta"

The meta package provides some simple convenience functions for dealing with
FASTQ files.

## Usage

#### func  CountNucleotides

```go
func CountNucleotides(r io.Reader) (int, error)
```
CountNucleotides returns the number of Nucleotides in the passed FASTQ formatted
Reader.

NB: No validation is performed on the data being read.

#### func  CountSequences

```go
func CountSequences(r io.Reader) (int, error)
```
CountSequences returns the number of Sequences in the passed FASTQ formatted
Reader.

NB: This function acts naively, simply counting the non-blank lines in the
Reader, and diving by 4.
