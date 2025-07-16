#!/bin/bash

set -euf -o pipefail;
cd "$(dirname "$0")";

code=0;

for i in a b c d e; do
	expectingSequences="$(cat "../tests/$i.sequences")";
	expectingNucleotides="$(cat "../tests/$i.nucleotides")";
	for j in "" ".gz"; do
		sequences="$(bash ./idoneal -s "../tests/$i.fastq$j")";
		if [ "$sequences" != "$expectingSequences" ]; then
			echo "Invalid sequences for $i.fastq$j, expecting $expectingSequences, got $sequences" >&2;
			code=1;
		fi;
		nucleotides="$(bash ./idoneal -n "../tests/$i.fastq$j")";
		if [ "$nucleotides" != "$expectingNucleotides" ]; then
			echo "Invalid nucleotides for $i.fastq$j, expecting $expectingNucleotides, got $nucleotides" >&2;
			code=1;
		fi;
	done;
done;
if [ $code -eq 0 ]; then
	echo "All Tests Succeeded!" >&2;
fi;

exit $code;
