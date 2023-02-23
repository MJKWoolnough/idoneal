// The tests package provides easy access to test data for testing.
package tests

import "embed"

// FS contains all of the embedded test data.
//go:embed *.fastq *.sequences
var FS embed.FS
