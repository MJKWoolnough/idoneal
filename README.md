# Idoneal

This project contains a CLI program to extract basic metadata from [FASTQ](https://en.wikipedia.org/wiki/FASTQ_format) formatted files.

## Installation

The following command can be used to install this program:

```
go install vimagination.zapto.org/idoneal/cmd/idoneal@v1.0.0
```

This will download and compile a `idoneal` binary into either your $GOPATH/bin or $GOBIN directory.

NB: You will need to have the [Go Programming Language](https://go.dev/) installed in order to use the above command.

## Command Line Flags

Usage: idoneal [OPTIONS] FILE

FILE should be in FASTQ format, and can be gzip compressed if the filename ends with .gz

|  Flag              |  Description                                 |
|--------------------|----------------------------------------------|
| -s / --sequences   | Print the number of sequences in the FILE.   |
| -n / --nucleotides | Print the number of nucleotides in the FILE. |


## Contributing

Pull Requests are accepted, as long as the following requirements are met:

 - All Go Code in Pull Requests should be formatted as per [gofmt](https://pkg.go.dev/cmd/gofmt).
 - All tests should pass.
 - If a new feature is being added, there should also be tests of the new functionality.
 - Likewise, if it is a bug fix, there should be a test against the old bad code that the new code fixes.
