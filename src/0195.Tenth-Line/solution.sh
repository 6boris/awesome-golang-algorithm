#!/usr/bin/env bash

cat README.md | awk '{if (NR==10){print $0}}'

# Read from the file file.txt and output the tenth line to stdout.
sed -n '10p' < file.txt