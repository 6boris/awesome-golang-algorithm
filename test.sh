#!/usr/bin/env bash

set -e
echo "" > coverage.txt

echo "test dir " $1

for d in $(go list ./$1... | grep -v vendor); do
    echo $d
    go test -json -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

