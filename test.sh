#!/bin/bash
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

test_dir="lcof"
echo "Test Dir ${test_dir}"
go clean -testcache
go test -coverprofile=profile.out -covermode=atomic -parallel 16 -p 16 ./$test_dir/*

test_dir="leetcode"
echo "Test Dir ${test_dir}"
for d in $(ls $test_dir); do
  {
    go test -coverprofile=profile.out -covermode=atomic -parallel 16 -p 16 ./$test_dir/$d/*
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
  } &
done
wait
