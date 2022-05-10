#!/bin/bash
for i in 1 2 3 4 5; do
  go run entrypoints/loadtest/main.go --type $1 > report-$1.csv && mv report-$1.csv results/m7-$1-$i.csv
done
